package logic

import (
	"Scheduler_go/common/workflow/engine"
	"Scheduler_go/service/workflow/model"
	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"reflect"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcInstRevokeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcInstRevokeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstRevokeLogic {
	return &ProcInstRevokeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcInstRevokeLogic) ProcInstRevoke(in *workflowclient.ProcInstRevokeReq) (*workflowclient.CommonResp, error) {
	resProcInst, err := l.svcCtx.ProcInstModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcInst没有该ID：%v", in.Id)
		}
		return nil, err
	}
	if in.Force == 1 {
		ProcExecution, err := l.svcCtx.ProcExecutionModel.FindOne(l.ctx, in.Id)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				return nil, fmt.Errorf("ProcExecution没有该ID：%v", in.Id)
			}
			return nil, err
		}
		if ProcExecution.NodeId != resProcInst.CurrentNodeId {
			return nil, fmt.Errorf("当前流程所在节点不是发起节点，无法撤销!")
		} else {

			resProcInst.Status = 2
			resProcInst.UpdatedName.String = in.UpdatedName
			resProcInst.UpdatedName.Valid = true
			resProcInst.UpdatedAt.Time = time.Now()
			resProcInst.UpdatedAt.Valid = true

			err = l.svcCtx.ProcInstModel.Update(l.ctx, resProcInst)

		}
	}

	//-----------------------------执行流程撤销事件 start-----------------------------

	//流程定义

	resProcDef, err := l.svcCtx.ProcDefModel.FindOne(l.ctx, resProcInst.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcDef没有该ID：%v", in.Id)
		}
		return nil, err
	}
	process, err := engine.ProcessParse(resProcDef.Resource)

	for _, e := range process.RevokeEvents {
		//log.Printf("正在处理节点[%s]中事件[%s]", CurrentNode.NodeName, e)
		//判断是否可以在事件池中获取事件
		event, ok := EventPool[e]
		if !ok {
			return nil, fmt.Errorf("事件%s未注册", e)
		}

		//拼装参数
		arg := []reflect.Value{
			reflect.ValueOf(event.S),
			reflect.ValueOf(resProcDef.Id),
			reflect.ValueOf(in.RevokeUserId),
		}

		//运行func
		result := event.M.Func.Call(arg)

		//如果选项IgnoreEventError为false,则说明需要验证事件是否出错
		//判断第一个返回参数是否为nil,若不是，则说明事件出错
		if !result[0].IsNil() {
			return nil, fmt.Errorf("流程[%s]撤销事件[%s]执行出错:%v", resProcDef.Name, event.M.Name, result[0])
		}
	}

	//-----------------------------执行流程撤销事件 end-----------------------------

	//开启事务
	// 开启事务添加
	err = l.svcCtx.ProcInstModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//将task表中所有该流程未finish的设置为finish

		resProcTask, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, in.Id)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				return fmt.Errorf("ProcTask没有该ID：%v", in.Id)
			}
			return err
		}

		resProcTask.UpdatedName.String = in.UpdatedName
		resProcTask.UpdatedName.Valid = true
		resProcTask.UpdatedAt.Time = time.Now()
		resProcTask.UpdatedAt.Valid = true
		resProcTask.IsFinished = 1
		resProcTask.FinishedTime = time.Now()

		err = l.svcCtx.ProcTaskModel.TransUpdate(ctx, session, resProcTask)

		//将task表中任务归档

		_, err = l.svcCtx.HistProcTaskModel.TransInsert(ctx, session, &model.HistProcTask{
			CreatedAt:          time.Now(),                     // 创建时间
			TaskId:             resProcTask.Id,                 // 任务ID
			ProcId:             resProcTask.ProcId,             // 流程ID
			ProcInstId:         resProcTask.ProcInstId,         // 流程实例ID
			BusinessId:         resProcTask.BusinessId,         // 业务ID
			Starter:            resProcTask.Starter,            // 流程发起人用户ID
			NodeId:             resProcTask.NodeId,             // 节点ID
			NodeName:           resProcTask.NodeName,           // 节点名称
			PrevNodeId:         resProcTask.PrevNodeId,         // 上个处理节点ID
			IsCosigned:         resProcTask.IsCosigned,         // 任意一人通过即可 1:会签
			BatchCode:          resProcTask.BatchCode,          // 批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别\"
			UserId:             resProcTask.UserId,             // 分配用户ID
			Status:             resProcTask.Status,             // 任务状态:0:初始 1:通过 2:驳回
			IsFinished:         resProcTask.IsFinished,         // 0:任务未完成 1:处理完成
			Comment:            resProcTask.Comment,            // 任务备注
			ProcInstCreateTime: resProcTask.ProcInstCreateTime, // 流程实例创建时间
			FinishedTime:       resProcTask.FinishedTime,       // 处理任务时间
			TenantId:           resProcTask.TenantId,           // 租户ID
			Data:               resProcTask.Data,               //
			CreatedName:        resProcTask.CreatedName,        // 创建人
		})
		if err != nil {
			return err
		}

		//删除task表中历史数据
		resProcTaskA, err := l.svcCtx.ProcTaskModel.FindOneByProcInstId(l.ctx, in.Id)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				return fmt.Errorf("ProcTask没有该ID：%v", in.Id)
			}
			return err
		}

		resProcTaskA.DeletedAt.Time = time.Now()
		resProcTaskA.DeletedAt.Valid = true
		resProcTaskA.DeletedName.String = in.UpdatedName
		resProcTaskA.DeletedName.Valid = true

		err = l.svcCtx.ProcTaskModel.TransUpdate(ctx, session, resProcTaskA)
		if err != nil {
			return err
		}
		//更新proc_inst表中状态

		resProcInst.Status = 2
		err = l.svcCtx.ProcInstModel.TransUpdate(ctx, session, resProcInst)

		if err != nil {
			return err
		}
		//将proc_inst表中数据归档
		_, err = l.svcCtx.HistProcInstModel.TransInsert(l.ctx, session, &model.HistProcInst{
			CreatedAt:     time.Now(),                                            // 创建时间
			ProcInstId:    sql.NullInt64{Int64: resProcInst.Id, Valid: true},     // 流程实例ID
			ProcId:        resProcInst.ProcId,                                    // 流程ID
			ProcName:      resProcInst.ProcName,                                  // 流程名称
			ProcVersion:   resProcInst.ProcVersion,                               // 流程版本号
			BusinessId:    resProcInst.BusinessId,                                // 业务ID
			Starter:       resProcInst.Starter,                                   // 流程发起人用户ID
			CurrentNodeId: resProcInst.CurrentNodeId,                             // 当前进行节点ID
			VariablesJson: resProcInst.VariablesJson,                             // 变量(Json)
			Status:        resProcInst.Status,                                    // 状态 0 未完成（审批中） 1 已完成 2 撤销
			TenantId:      in.TenantId,                                           // 租户ID
			Data:          sql.NullString{String: in.Data, Valid: in.Data != ""}, //
			CreatedName:   resProcInst.CreatedName,                               // 创建人
		})
		if err != nil {
			return err
		}

		//删除proc_inst表中历史数据
		resProcInst.DeletedAt.Time = time.Now()
		resProcInst.DeletedAt.Valid = true
		resProcInst.DeletedName.String = resProcInst.CreatedName
		resProcInst.DeletedName.Valid = true

		err = l.svcCtx.ProcInstModel.Update(l.ctx, resProcInst)
		if err != nil {
			return err
		}

		//将proc_inst_variable表中数据归档
		resProcInstVariable, err := l.svcCtx.ProcInstVariableModel.FindOne(l.ctx, in.Id)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				return fmt.Errorf("ProcInstVariable没有该ID：%v", in.Id)
			}
			return err
		}
		_, err = l.svcCtx.HistProcInstVariableModel.Insert(l.ctx, &model.HistProcInstVariable{
			CreatedAt:   time.Now(),                                            // 创建时间
			ProcInstId:  resProcInstVariable.Id,                                // 流程实例ID
			Key:         resProcInstVariable.Key,                               // 变量key
			Value:       resProcInstVariable.Value,                             // 变量value
			TenantId:    in.TenantId,                                           // 租户ID
			Data:        sql.NullString{String: in.Data, Valid: in.Data != ""}, //
			CreatedName: resProcInstVariable.CreatedName,                       // 创建人
		})
		if err != nil {
			return err
		}

		//删除proc_inst_variable表中历史数据

		resProcInstVariable.DeletedAt.Time = time.Now()
		resProcInstVariable.DeletedAt.Valid = true
		resProcInstVariable.DeletedName.String = resProcInstVariable.CreatedName
		resProcInstVariable.DeletedName.Valid = true

		err = l.svcCtx.ProcInstVariableModel.Update(l.ctx, resProcInstVariable)
		if err != nil {
			return err
		}

		return nil
	})

	return &workflowclient.CommonResp{}, nil

}

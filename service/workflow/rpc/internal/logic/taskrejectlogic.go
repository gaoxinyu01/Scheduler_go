package logic

import (
	"Scheduler_go/common/workflow/modelx"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskRejectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskRejectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskRejectLogic {
	return &TaskRejectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 任务驳回
func (l *TaskRejectLogic) TaskReject(in *workflowclient.TaskRejectReq) (*workflowclient.CommonResp, error) {
	////获取节点信息
	taskInfo, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcTask没有该ID：%v", in.TaskId)
		}
		return nil, err
	}
	//获取task所在的node

	resProcInst, err := l.svcCtx.ProcInstModel.FindOne(l.ctx, taskInfo.ProcInstId)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcInst没有该ID：%v", taskInfo.ProcId)
		}
		return nil, err
	}

	//从Cache中获得流程节点列表
	Nodes, err := GetProcCache(l.ctx, l.svcCtx.ProcDefModel, resProcInst.ProcId)
	if err != nil {
		return nil, err
	}
	//获得节点
	CurrentNode, ok := Nodes[taskInfo.NodeId]
	if !ok {
		return nil, fmt.Errorf("ID为%d的流程实例中不存在ID为%s的节点", resProcInst.ProcId, taskInfo.NodeId)
	}

	//判断节点是否已处理
	if taskInfo.IsFinished == 1 {
		return nil, fmt.Errorf("节点ID%d已处理，无需操作", taskInfo.Id)
	}
	//------------------------如果是通过，且DirectlyToWhoRejectedMe为true,则需做功能前置验证 ------------------------
	//1、是否是会签节点
	//2、是否存在上一个任务节点?上一个节点是否做的是驳回

	//会签节点无法使用此功能，因为会签节点没有“统一意志”
	if taskInfo.IsCosigned == 1 {
		return nil, errors.New("会签节点无法使用【DirectlyToWhoRejectedMe】功能!")
	}

	//任务没有上级节点
	if taskInfo.PrevNodeId == "" {
		return nil, errors.New("此任务不存在上级节点,无法使用【DirectlyToWhoRejectedMe】功能!!")
	}

	//判断任务的上一个节点是不是做了驳回

	whereBuilder := l.svcCtx.ProcTaskModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"batch_code ": taskInfo.BatchCode,
	})
	// 任务状态:0:初始 1:通过 2:驳回
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"status ": 2,
	})

	_, err = l.svcCtx.ProcTaskModel.FindList(l.ctx, whereBuilder, 1, 999)
	if err != nil {
		return nil, errors.New("此任务的上一节点并未做驳回,无法使用【DirectlyToWhoRejectedMe】功能！")
	}

	//如果是驳回，则验证是否起始节点(起始节点不能做驳回)
	if CurrentNode.NodeType == modelx.RootNode {
		return nil, errors.New("起始节点无法驳回!")
	}
	//将任务提交数据(通过、驳回、变量)保存到数据库
	//判断节点是否已处理
	if taskInfo.IsFinished == 1 {
		return nil, fmt.Errorf("节点ID%d已处理，无需操作", taskInfo.Id)
	}
	// 开启事务添加
	err = l.svcCtx.ProcTaskModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {
		//更新task表记录
		taskInfo.UpdatedName.String = in.CreatedName
		taskInfo.UpdatedName.Valid = true
		taskInfo.UpdatedAt.Time = time.Now()
		taskInfo.UpdatedAt.Valid = true
		taskInfo.Status = 2
		taskInfo.IsFinished = 1
		taskInfo.Comment = sql.NullString{String: in.Comment, Valid: in.Comment != ""}
		taskInfo.FinishedTime = time.Now()

		err = l.svcCtx.ProcTaskModel.TransUpdate(ctx, sqlx, taskInfo)
		if err != nil {
			return err
		}

		//1、非会签节点，一人通过即通过，所以要把其他人的任务finish掉
		//2、不论是否会签，都是一人驳回即驳回，所以需要把同一批次task的isfinish设置为1,让其他人不用再处理
		if (taskInfo.IsCosigned == 0 && taskInfo.Status == 1) || taskInfo.Status == 2 {
			taskInfo.UpdatedName.String = in.CreatedName
			taskInfo.UpdatedName.Valid = true
			taskInfo.UpdatedAt.Time = time.Now()
			taskInfo.UpdatedAt.Valid = true
			taskInfo.IsFinished = 1
			taskInfo.FinishedTime = time.Now()
		}
		//设置实例变量

		return nil
	})

	return &workflowclient.CommonResp{}, nil
}

package logic

import (
	"Scheduler_go/common/workflow/engine"
	"Scheduler_go/common/workflow/modelx"
	"Scheduler_go/service/workflow/model"
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

type ProcDefSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcDefSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcDefSaveLogic {
	return &ProcDefSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流程定义
func (l *ProcDefSaveLogic) ProcDefSave(in *workflowclient.ProcDefSaveReq) (resp *workflowclient.CommonResp, err error) {

	//解析传入的json，获得process数据结构
	process, err := engine.ProcessParse(in.Resource)
	if err != nil {
		return nil, err
	}

	if process.ProcessName == "" || process.Source == "" || in.CreateUserId == "" {
		return nil, fmt.Errorf("流程名称、来源、创建人ID不能为空")
	}
	//首先判断此工作流是否已定义

	whereBuilder := l.svcCtx.ProcDefModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 流程名称
	if len(process.ProcessName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"name ": process.ProcessName,
		})
	}

	// 来源
	if len(process.Source) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"source ": process.Source,
		})
	}

	allProcDef, err := l.svcCtx.ProcDefModel.FindList(l.ctx, whereBuilder, 0, 99999999)
	if err != nil {
		return nil, err
	}

	// 开启事务添加
	err = l.svcCtx.ProcDefModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {
		//判断工作流是否已经定义

		for _, res := range allProcDef {
			if res.Id != 0 { //已有老版本
				//将老版本移到历史表中
				_, err = l.svcCtx.HistProcDefModel.TransInsert(ctx, sqlx, &model.HistProcDef{
					CreatedAt:    time.Now(),       // 创建时间
					ProcId:       res.Id,           // 流程模板ID
					Name:         res.Name,         // 流程名称
					Version:      res.Version,      // 版本号
					ProcType:     res.ProcType,     // 流程类型
					Resource:     res.Resource,     // 流程定义模板
					CreateUserId: res.CreateUserId, // 创建者ID
					Source:       res.Source,       // 来源
					TenantId:     res.TenantId,     // 租户ID
					Data:         res.Data,         //
					CreatedName:  res.CreatedName,  // 创建人
				})
				if err != nil {
					return err
				}
				//更新现有定义
				res.Version = res.Version + 1
				res.UpdatedAt.Time = time.Now()
				res.UpdatedAt.Valid = true
				err = l.svcCtx.ProcDefModel.Update(l.ctx, res)

			} else {
				//若没有老版本，则直接插入
				_, err = l.svcCtx.ProcDefModel.TransInsert(ctx, sqlx, &model.ProcDef{
					CreatedAt:    time.Now(),                                                            // 创建时间
					Name:         process.ProcessName,                                                   // 流程名称
					Version:      1,                                                                     // 版本号
					ProcType:     1,                                                                     // 流程类型
					Resource:     in.Resource,                                                           // 流程定义模板
					CreateUserId: sql.NullString{String: in.CreateUserId, Valid: in.CreateUserId != ""}, // 创建者ID
					Source:       sql.NullString{String: process.Source, Valid: process.Source != ""},   // 来源
					TenantId:     in.TenantId,                                                           // 租户ID
					Data:         sql.NullString{String: in.Data, Valid: in.Data != ""},                 //
					CreatedName:  in.CreatedName,                                                        // 创建人
				})
				if err != nil {
					return err
				}
			}

			//将proc_execution表对应数据移到历史表中
			_, err = l.svcCtx.HistProcExecutionModel.TransInsert(ctx, sqlx, &model.HistProcExecution{
				CreatedAt:   time.Now(),                                            // 创建时间
				ProcId:      res.Id,                                                // 实例ID
				ProcVersion: res.Version,                                           // 流程版本号
				ProcName:    res.Name,                                              // 流程名
				NodeId:      process.Nodes[].NodeID,                                // 节点ID
				NodeName:    process.Nodes[].NodeName,                              // 节点名称
				PrevNodeId:  process.Nodes[].PrevNodeIDs[],                         // 上级节点ID
				NodeType:    int64(process.Nodes[].NodeType),                       // 节点类型 0 开始节点，1 任务节点 ，2 网关节点，3 结束节点
				IsCosigned:  int64((process.Nodes[].IsCosigned)),                   // 是否会签  0 不会签  1 会签
				TenantId:    in.TenantId,                                           // 租户ID
				Data:        sql.NullString{String: in.Data, Valid: in.Data != ""}, //
				CreatedName: in.CreatedName,                                        // 创建人
			})
			if err != nil {
				return err
			}
			//删除proc_execution表对应数据
			resProcExecution, err := l.svcCtx.ProcExecutionModel.FindOneByProcId(l.ctx, res.Id)
			if err != nil {
				if errors.Is(err, sqlc.ErrNotFound) {
					return fmt.Errorf("ProcExecution没有该ID：%v", res.Id)
				}
				return err
			}

			// 判断该数据是否被删除
			if res.DeletedAt.Valid == true {
				return fmt.Errorf("ProcExecution该ID已被删除：%v", res.Id)
			}
			if res.TenantId != in.TenantId {
				return errors.New("不是一个租户非法操作")
			}

			resProcExecution.DeletedAt.Time = time.Now()
			resProcExecution.DeletedAt.Valid = true
			resProcExecution.DeletedName.String = in.CreatedName
			resProcExecution.DeletedName.Valid = true

			err = l.svcCtx.ProcExecutionModel.TransUpdate(ctx, sqlx, resProcExecution)
			if err != nil {
				return err
			}
			//解析node之间的关系，流程节点执行关系定义记录
			execution := nodes2Execution(res.Id, res.Version, process.Nodes)
			//将Execution定义插入proc_execution表
			for _, proExecution := range execution {
				_, err = l.svcCtx.ProcExecutionModel.TransInsert(ctx, sqlx, &model.ProcExecution{
					CreatedAt:   time.Now(),                                                                // 创建时间
					ProcId:      proExecution.ProcId,                                                       // 实例ID
					ProcVersion: proExecution.ProcVersion,                                                  // 流程版本号
					ProcName:    proExecution.ProcName,                                                     // 流程名
					NodeId:      proExecution.NodeId,                                                       // 节点ID
					NodeName:    proExecution.NodeName,                                                     // 节点名称
					PrevNodeId:  proExecution.PrevNodeId,                                                   // 上级节点ID
					NodeType:    proExecution.NodeType,                                                     // 节点类型 0 开始节点，1 任务节点 ，2 网关节点，3 结束节点
					IsCosigned:  proExecution.IsCosigned,                                                   // 是否会签  0 不会签  1 会签
					TenantId:    in.TenantId,                                                               // 租户ID
					Data:        sql.NullString{String: proExecution.Data, Valid: proExecution.Data != ""}, //
					CreatedName: proExecution.CreatedName,                                                  // 创建人
				})
				if err != nil {
					return err
				}
			}

		}

		return nil
	})

	return &workflowclient.CommonResp{}, nil

}
func nodes2Execution(ProcID int64, ProcVersion int64, nodes []modelx.Node) []workflowclient.ProcExecutionListData {
	var executions []workflowclient.ProcExecutionListData
	for _, n := range nodes {
		if len(n.PrevNodeIDs) <= 1 { //上级节点数<=1的情况下
			var PrevNodeID string
			if len(n.PrevNodeIDs) == 0 { //开始节点没有上级
				PrevNodeID = ""
			} else {
				PrevNodeID = n.PrevNodeIDs[0]
			}
			executions = append(executions, workflowclient.ProcExecutionListData{
				ProcId:      ProcID,
				ProcVersion: ProcVersion,
				NodeId:      n.NodeID,
				NodeName:    n.NodeName,
				PrevNodeId:  PrevNodeID,
				NodeType:    int64(n.NodeType),
				IsCosigned:  int64(n.IsCosigned),
				CreatedAt:   time.Now().UnixMilli(),
			})
		} else { //上级节点>1的情况下，则每一个上级节点都要生成一行
			for _, prev := range n.PrevNodeIDs {
				executions = append(executions, workflowclient.ProcExecutionListData{
					ProcId:      ProcID,
					ProcVersion: ProcVersion,
					NodeId:      n.NodeID,
					NodeName:    n.NodeName,
					PrevNodeId:  prev,
					NodeType:    int64(n.NodeType),
					IsCosigned:  int64(n.IsCosigned),
					CreatedAt:   time.Now().UnixMilli(),
				})
			}
		}
	}
	return executions
}

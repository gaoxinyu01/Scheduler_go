package logic

import (
	"Scheduler_go/service/workflow/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskTransferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskTransferLogic {
	return &TaskTransferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 将任务转交给他人处理
func (l *TaskTransferLogic) TaskTransfer(in *workflowclient.TaskTransferReq) (*workflowclient.CommonResp, error) {
	//传入用户去重
	users := MakeUnique(in.Users)
	if len(users) < 1 {
		return nil, errors.New("转让任务操作必须指定至少一个候选人")
	}

	//获得task信息
	taskInfo, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcTask没有该ID：%v", in.TaskId)
		}
		return nil, err
	}

	//已完成任务不能转交
	if taskInfo.IsFinished == 1 {
		return nil, errors.New("任务已完成，无法转交")
	}

	//思考：考虑会签情况下，如果仅仅设置为结束，则节点由于永远不会达到"任务总数与通过总人数一致"，成为死节点
	//所以，转交任务需将原任务删除

	//开启事务

	// 开启事务添加
	err = l.svcCtx.ProcTaskModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//首先删除任务

		taskInfo.DeletedAt.Time = time.Now()
		taskInfo.DeletedAt.Valid = true
		taskInfo.DeletedName.String = in.CreatedName
		taskInfo.DeletedName.Valid = true

		err = l.svcCtx.ProcTaskModel.TransUpdate(l.ctx, session, taskInfo)
		if err != nil {
			return err
		}

		//生成新任务数据
		for _, u := range users {
			_, err = l.svcCtx.ProcTaskModel.TransInsert(l.ctx, session, &model.ProcTask{
				CreatedAt:          time.Now(),                                            // 创建时间
				ProcId:             taskInfo.ProcId,                                       // 流程ID
				ProcInstId:         taskInfo.ProcInstId,                                   // 流程实例ID
				BusinessId:         taskInfo.BusinessId,                                   // 业务ID
				Starter:            taskInfo.Starter,                                      // 流程发起人用户ID
				NodeId:             taskInfo.NodeId,                                       // 节点ID
				NodeName:           taskInfo.NodeName,                                     // 节点名称
				PrevNodeId:         taskInfo.PrevNodeId,                                   // 上个处理节点ID
				IsCosigned:         taskInfo.IsCosigned,                                   // 任意一人通过即可 1:会签
				BatchCode:          taskInfo.BatchCode,                                    // 批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别\"
				UserId:             u,                                                     // 分配用户ID
				Status:             taskInfo.Status,                                       // 任务状态:0:初始 1:通过 2:驳回
				IsFinished:         taskInfo.IsFinished,                                   // 0:任务未完成 1:处理完成
				Comment:            taskInfo.Comment,                                      // 任务备注
				ProcInstCreateTime: taskInfo.ProcInstCreateTime,                           // 流程实例创建时间
				FinishedTime:       taskInfo.FinishedTime,                                 // 处理任务时间
				TenantId:           in.TenantId,                                           // 租户ID
				Data:               sql.NullString{String: in.Data, Valid: in.Data != ""}, //
				CreatedName:        in.CreatedName,                                        // 创建人
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return &workflowclient.CommonResp{}, nil
}

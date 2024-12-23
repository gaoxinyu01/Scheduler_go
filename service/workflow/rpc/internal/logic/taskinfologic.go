package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInfoLogic {
	return &TaskInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 任务信息
func (l *TaskInfoLogic) TaskInfo(in *workflowclient.TaskInfoReq) (*workflowclient.TaskInfoResp, error) {
	res, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, in.Taskid)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcTask没有该ID：%v", in.Taskid)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("ProcTask该ID已被删除：%v", in.Taskid)
	}
	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	return &workflowclient.TaskInfoResp{
		Id:                 res.Id,                                  //任务ID
		ProcId:             res.ProcId,                              //流程ID
		ProcInstId:         res.ProcInstId,                          //流程实例ID
		BusinessId:         res.BusinessId,                          //业务ID
		Starter:            res.Starter,                             //流程发起人用户ID
		NodeId:             res.NodeId,                              //节点ID
		NodeName:           res.NodeName,                            //节点名称
		PrevNodeId:         res.PrevNodeId,                          //上个处理节点ID
		IsCosigned:         res.IsCosigned,                          //任意一人通过即可 1:会签
		BatchCode:          res.BatchCode,                           //批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别\"
		UserId:             res.UserId,                              //分配用户ID
		Status:             res.Status,                              //任务状态:0:初始 1:通过 2:驳回
		IsFinished:         res.IsFinished,                          //0:任务未完成 1:处理完成
		Comment:            res.Comment.String,                      //任务备注
		ProcInstCreateTime: res.ProcInstCreateTime.Time.UnixMilli(), //流程实例创建时间
		FinishedTime:       res.FinishedTime.UnixMilli(),            //处理任务时间
		Data:               res.Data.String,                         //
		CreatedAt:          res.CreatedAt.UnixMilli(),               //创建时间
		UpdatedAt:          res.UpdatedAt.Time.UnixMilli(),          //更新时间
		CreatedName:        res.CreatedName,                         //创建人
		UpdatedName:        res.UpdatedName.String,                  //更新人
	}, nil
}

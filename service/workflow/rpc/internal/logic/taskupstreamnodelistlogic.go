package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskUpstreamNodeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskUpstreamNodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskUpstreamNodeListLogic {
	return &TaskUpstreamNodeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取本任务所在节点的所有上游节点
func (l *TaskUpstreamNodeListLogic) TaskUpstreamNodeList(in *workflowclient.TaskUpstreamNodeListReq) (*workflowclient.TaskUpstreamNodeListResp, error) {
	//获得task信息
	task, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, in.Taskid)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcTask没有该ID：%v", in.Taskid)
		}
		return nil, err
	}

	whereBuilder := l.svcCtx.ProcExecutionModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 节点ID
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"node_id ": task.NodeId,
	})

	all, err := l.svcCtx.ProcExecutionModel.FindList(l.ctx, whereBuilder, 1, 9999)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.ProcExecutionModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	countBuilder = countBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	countBuilder = countBuilder.Where(squirrel.Eq{
		"node_id ": task.NodeId,
	})
	count, err := l.svcCtx.ProcExecutionModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*workflowclient.TaskUpstreamNodeListDate
	for _, item := range all {
		list = append(list, &workflowclient.TaskUpstreamNodeListDate{
			Id:          item.Id,     //执行ID
			ProcId:      item.ProcId, //实例ID
			ProcInstId:  item.ProcId,
			NodeId:      item.NodeId,                     //节点ID
			NodeName:    item.NodeName,                   //节点名称
			PrevNodeId:  item.PrevNodeId,                 //上级节点ID
			IsCosigned:  item.IsCosigned,                 //是否会签  0 不会签  1 会签
			Data:        item.Data.String,                //
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName: item.CreatedName,                //创建人
			UpdatedName: item.UpdatedName.String,         //更新人
		})
	}

	return &workflowclient.TaskUpstreamNodeListResp{
		Total: count,
		List:  list,
	}, nil

}

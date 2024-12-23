package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcExecutionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcExecutionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcExecutionListLogic {
	return &ProcExecutionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 执行任务
func (l *ProcExecutionListLogic) ProcExecutionList(in *workflowclient.ProcExecutionListReq) (resp *workflowclient.ProcExecutionListResp, err error) {

	whereBuilder := l.svcCtx.ProcExecutionModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 实例ID
	if in.ProcId != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"proc_id ": in.ProcId,
		})
	}
	// 流程版本号
	if in.ProcVersion != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"proc_version ": in.ProcVersion,
		})
	}
	// 流程名
	if len(in.ProcName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"proc_name ": "%" + in.ProcName + "%",
		})
	}
	// 节点ID
	if len(in.NodeId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"node_id ": "%" + in.NodeId + "%",
		})
	}
	// 节点名称
	if len(in.NodeName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"node_name ": "%" + in.NodeName + "%",
		})
	}
	// 上级节点ID
	if len(in.PrevNodeId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"prev_node_id ": "%" + in.PrevNodeId + "%",
		})
	}
	// 节点类型 0 开始节点，1 任务节点 ，2 网关节点，3 结束节点
	if in.NodeType != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"node_type ": in.NodeType,
		})
	}
	// 是否会签  0 不会签  1 会签
	if in.IsCosigned != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"is_cosigned ": in.IsCosigned,
		})
	}
	//
	if len(in.Data) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"data ": "%" + in.Data + "%",
		})
	}

	all, err := l.svcCtx.ProcExecutionModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.ProcExecutionModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	countBuilder = countBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 实例ID
	if in.ProcId != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"proc_id ": in.ProcId,
		})
	}
	// 流程版本号
	if in.ProcVersion != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"proc_version ": in.ProcVersion,
		})
	}
	// 流程名
	if len(in.ProcName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"proc_name ": "%" + in.ProcName + "%",
		})
	}
	// 节点ID
	if len(in.NodeId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"node_id ": "%" + in.NodeId + "%",
		})
	}
	// 节点名称
	if len(in.NodeName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"node_name ": "%" + in.NodeName + "%",
		})
	}
	// 上级节点ID
	if len(in.PrevNodeId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"prev_node_id ": "%" + in.PrevNodeId + "%",
		})
	}
	// 节点类型 0 开始节点，1 任务节点 ，2 网关节点，3 结束节点
	if in.NodeType != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"node_type ": in.NodeType,
		})
	}
	// 是否会签  0 不会签  1 会签
	if in.IsCosigned != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"is_cosigned ": in.IsCosigned,
		})
	}
	//
	if len(in.Data) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"data ": "%" + in.Data + "%",
		})
	}
	count, err := l.svcCtx.ProcExecutionModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*workflowclient.ProcExecutionListData
	for _, item := range all {
		list = append(list, &workflowclient.ProcExecutionListData{
			Id:          item.Id,                         //执行ID
			ProcId:      item.ProcId,                     //实例ID
			ProcVersion: item.ProcVersion,                //流程版本号
			ProcName:    item.ProcName,                   //流程名
			NodeId:      item.NodeId,                     //节点ID
			NodeName:    item.NodeName,                   //节点名称
			PrevNodeId:  item.PrevNodeId,                 //上级节点ID
			NodeType:    item.NodeType,                   //节点类型 0 开始节点，1 任务节点 ，2 网关节点，3 结束节点
			IsCosigned:  item.IsCosigned,                 //是否会签  0 不会签  1 会签
			Data:        item.Data.String,                //
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName: item.CreatedName,                //创建人
			UpdatedName: item.UpdatedName.String,         //更新人
		})
	}

	return &workflowclient.ProcExecutionListResp{
		Total: count,
		List:  list,
	}, nil
}

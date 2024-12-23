package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcInstListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcInstListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstListLogic {
	return &ProcInstListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcInstListLogic) ProcInstList(in *workflowclient.ProcInstListReq) (resp *workflowclient.ProcInstListResp, err error) {

	whereBuilder := l.svcCtx.ProcInstModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	if len(in.TenantId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"tenant_id": in.TenantId,
		})
	}
	// 流程ID
	if in.ProcId != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"proc_id ": in.ProcId,
		})
	}
	// 流程名称
	if len(in.ProcName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"proc_name ": "%" + in.ProcName + "%",
		})
	}
	// 流程版本号
	if in.ProcVersion != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"proc_version ": in.ProcVersion,
		})
	}
	// 业务ID
	if len(in.BusinessId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"business_id ": "%" + in.BusinessId + "%",
		})
	}
	// 流程发起人用户ID
	if len(in.Starter) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"starter ": "%" + in.Starter + "%",
		})
	}
	// 当前进行节点ID
	if len(in.CurrentNodeId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"current_node_id ": "%" + in.CurrentNodeId + "%",
		})
	}
	// 变量(Json)
	if len(in.VariablesJson) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"variables_json ": "%" + in.VariablesJson + "%",
		})
	}
	// 状态 0 未完成（审批中） 1 已完成 2 撤销
	if in.Status != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"status ": in.Status,
		})
	}
	//
	if len(in.Data) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"data ": "%" + in.Data + "%",
		})
	}

	all, err := l.svcCtx.ProcInstModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.ProcInstModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")
	if len(in.TenantId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"tenant_id": in.TenantId,
		})
	}
	// 流程ID
	if in.ProcId != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"proc_id ": in.ProcId,
		})
	}
	// 流程名称
	if len(in.ProcName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"proc_name ": "%" + in.ProcName + "%",
		})
	}
	// 流程版本号
	if in.ProcVersion != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"proc_version ": in.ProcVersion,
		})
	}
	// 业务ID
	if len(in.BusinessId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"business_id ": "%" + in.BusinessId + "%",
		})
	}
	// 流程发起人用户ID
	if len(in.Starter) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"starter ": "%" + in.Starter + "%",
		})
	}
	// 当前进行节点ID
	if len(in.CurrentNodeId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"current_node_id ": "%" + in.CurrentNodeId + "%",
		})
	}
	// 变量(Json)
	if len(in.VariablesJson) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"variables_json ": "%" + in.VariablesJson + "%",
		})
	}
	// 状态 0 未完成（审批中） 1 已完成 2 撤销
	if in.Status != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"status ": in.Status,
		})
	}
	//
	if len(in.Data) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"data ": "%" + in.Data + "%",
		})
	}
	count, err := l.svcCtx.ProcInstModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*workflowclient.ProcInstListData
	for _, item := range all {
		list = append(list, &workflowclient.ProcInstListData{
			Id:            item.Id,                         //流程实例ID
			ProcId:        item.ProcId,                     //流程ID
			ProcName:      item.ProcName,                   //流程名称
			ProcVersion:   item.ProcVersion,                //流程版本号
			BusinessId:    item.BusinessId,                 //业务ID
			Starter:       item.Starter,                    //流程发起人用户ID
			CurrentNodeId: item.CurrentNodeId,              //当前进行节点ID
			VariablesJson: item.VariablesJson.String,       //变量(Json)
			Status:        item.Status,                     //状态 0 未完成（审批中） 1 已完成 2 撤销
			Data:          item.Data.String,                //
			CreatedAt:     item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:     item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:   item.CreatedName,                //创建人
			UpdatedName:   item.UpdatedName.String,         //更新人
		})
	}

	return &workflowclient.ProcInstListResp{
		Total: count,
		List:  list,
	}, nil
}

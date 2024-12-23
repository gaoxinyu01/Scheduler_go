package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistProcDefListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHistProcDefListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistProcDefListLogic {
	return &HistProcDefListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流程定义历史
func (l *HistProcDefListLogic) HistProcDefList(in *workflowclient.HistProcDefListReq) (resp *workflowclient.HistProcDefListResp, err error) {

	whereBuilder := l.svcCtx.HistProcDefModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 流程模板ID
	if in.ProcId != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"proc_id ": in.ProcId,
		})
	}
	// 流程名称
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 版本号
	if in.Version != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"version ": in.Version,
		})
	}
	// 流程类型
	if in.ProcType != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"proc_type ": in.ProcType,
		})
	}
	// 流程定义模板
	if len(in.Resource) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"resource ": "%" + in.Resource + "%",
		})
	}
	// 创建者ID
	if len(in.CreateUserId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"create_user_id ": "%" + in.CreateUserId + "%",
		})
	}
	// 来源
	if len(in.Source) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"source ": "%" + in.Source + "%",
		})
	}
	//
	if len(in.Data) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"data ": "%" + in.Data + "%",
		})
	}

	all, err := l.svcCtx.HistProcDefModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.HistProcDefModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	countBuilder = countBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 流程模板ID
	if in.ProcId != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"proc_id ": in.ProcId,
		})
	}
	// 流程名称
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 版本号
	if in.Version != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"version ": in.Version,
		})
	}
	// 流程类型
	if in.ProcType != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"proc_type ": in.ProcType,
		})
	}
	// 流程定义模板
	if len(in.Resource) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"resource ": "%" + in.Resource + "%",
		})
	}
	// 创建者ID
	if len(in.CreateUserId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"create_user_id ": "%" + in.CreateUserId + "%",
		})
	}
	// 来源
	if len(in.Source) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"source ": "%" + in.Source + "%",
		})
	}
	//
	if len(in.Data) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"data ": "%" + in.Data + "%",
		})
	}
	count, err := l.svcCtx.HistProcDefModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*workflowclient.HistProcDefListData
	for _, item := range all {
		list = append(list, &workflowclient.HistProcDefListData{
			Id:           item.Id,                         //ID
			ProcId:       item.ProcId,                     //流程模板ID
			Name:         item.Name,                       //流程名称
			Version:      item.Version,                    //版本号
			ProcType:     item.ProcType,                   //流程类型
			Resource:     item.Resource,                   //流程定义模板
			CreateUserId: item.CreateUserId.String,        //创建者ID
			Source:       item.Source.String,              //来源
			Data:         item.Data.String,                //
			CreatedAt:    item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:    item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:  item.CreatedName,                //创建人
			UpdatedName:  item.UpdatedName.String,         //更新人
		})
	}

	return &workflowclient.HistProcDefListResp{
		Total: count,
		List:  list,
	}, nil
}
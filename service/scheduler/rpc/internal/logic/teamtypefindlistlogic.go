package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeamTypeFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTeamTypeFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamTypeFindListLogic {
	return &TeamTypeFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TeamTypeFindListLogic) TeamTypeFindList(in *schedulerclient.TeamTypeFindListReq) (resp *schedulerclient.TeamTypeFindListResp, err error) {

	whereBuilder := l.svcCtx.TeamTypeModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 部门名称
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 描述
	if len(in.Description) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"description ": "%" + in.Description + "%",
		})
	}

	all, err := l.svcCtx.TeamTypeModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.TeamTypeModel.CountBuilder("id")

	countBuilder = countBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 部门名称
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 描述
	if len(in.Description) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"description ": "%" + in.Description + "%",
		})
	}
	count, err := l.svcCtx.TeamTypeModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*schedulerclient.TeamTypeListData
	for _, item := range all {
		list = append(list, &schedulerclient.TeamTypeListData{
			Id:          item.Id,          //部门类型ID
			Name:        item.Name,        //部门名称
			Description: item.Description, //描述
		})
	}

	return &schedulerclient.TeamTypeFindListResp{
		Total: count,
		List:  list,
	}, nil
}

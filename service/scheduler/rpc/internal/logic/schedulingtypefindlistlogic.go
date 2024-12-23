package logic

import (
	"context"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingTypeFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSchedulingTypeFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingTypeFindListLogic {
	return &SchedulingTypeFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SchedulingTypeFindListLogic) SchedulingTypeFindList(in *schedulerclient.SchedulingTypeFindListReq) (*schedulerclient.SchedulingTypeFindListResp, error) {
	all, err := l.svcCtx.SchedulingTypeModel.FindListByTenantId(in.Current, in.PageSize, in.TenantId)
	if err != nil {
		return nil, err
	}

	var list []*schedulerclient.SchedulingTypeFindListData
	for _, item := range *all {
		list = append(list, &schedulerclient.SchedulingTypeFindListData{
			Id:        item.Id,
			Name:      item.Name,
			StartTime: item.StartTime,
			EndTime:   item.EndTime,
			Colour:    item.Colour.String,
			Remark:    item.Remark.String,
		})

	}

	total := l.svcCtx.SchedulingTypeModel.CountByTenantId(in.TenantId)

	return &schedulerclient.SchedulingTypeFindListResp{
		Total: total,
		List:  list,
	}, nil
}

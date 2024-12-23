package logic

import (
	"context"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSchedulingFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingFindListLogic {
	return &SchedulingFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SchedulingFindListLogic) SchedulingFindList(in *schedulerclient.SchedulingFindListReq) (*schedulerclient.SchedulingFindListResp, error) {
	all, err := l.svcCtx.SchedulingModel.FindList(in.Current, in.PageSize, in.Time, in.Name,
		in.StartTime, in.EndTime, in.TeamName, in.UserName, in.TenantId)
	if err != nil {
		return nil, err
	}

	var list []*schedulerclient.SchedulingFindListData
	for _, item := range *all {
		list = append(list, &schedulerclient.SchedulingFindListData{
			Id:           item.Id.String,
			CreatedAt:    item.CreatedAt.Time.UnixMilli(),
			UpdatedAt:    item.UpdatedAt.Time.UnixMilli(),
			CreatedName:  item.CreatedName.String,
			UpdatedName:  item.UpdatedName.String,
			Time:         item.Time.Int64,
			Name:         item.Name.String,
			StartTime:    item.StartTime.Int64,
			EndTime:      item.EndTime.Int64,
			Colour:       item.Colour.String,
			TeamName:     item.TeamName.String,
			UserName:     item.UserName.String,
			JobStartTime: item.JobStartTime.Int64,
			JobEndTime:   item.JobEndTime.Int64,
		})
	}

	total := l.svcCtx.SchedulingModel.Count(in.Time, in.Name, in.StartTime, in.EndTime, in.TeamName, in.UserName, in.TenantId)

	return &schedulerclient.SchedulingFindListResp{
		Total: total,
		List:  list,
	}, nil
}

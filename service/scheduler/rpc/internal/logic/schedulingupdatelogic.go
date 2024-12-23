package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSchedulingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingUpdateLogic {
	return &SchedulingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SchedulingUpdateLogic) SchedulingUpdate(in *schedulerclient.SchedulingUpdateReq) (*schedulerclient.CommonResp, error) {
	res, err := l.svcCtx.SchedulingModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("没有该数据")
		}
		return nil, err
	}

	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	if in.Time != 0 {
		res.Time = in.Time
	}
	if in.Name != "" {
		res.Name = in.Name
	}
	if in.StartTime != 0 {
		res.StartTime = in.StartTime
	}
	if in.EndTime != 0 {
		res.EndTime = in.EndTime
	}
	if in.Colour != "" {
		res.Colour.String = in.Colour
		res.Colour.Valid = in.Colour != ""
	}
	if in.UserName != "" {
		res.UserName = in.UserName
	}
	if in.TeamId != "" {
		resTeamType, err := l.svcCtx.TeamTypeModel.FindOne(l.ctx, in.TeamId)
		if err != nil {
			if err == sqlc.ErrNotFound {
				return nil, errors.New("没有该部门")
			}
			return nil, err
		}
		res.TeamId = in.TeamId
		res.TeamName = resTeamType.Name
	}
	if in.UserId != "" {
		res.UserId = in.UserId
	}
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true
	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true

	err = l.svcCtx.SchedulingModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

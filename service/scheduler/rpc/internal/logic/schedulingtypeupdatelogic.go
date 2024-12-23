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

type SchedulingTypeUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSchedulingTypeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingTypeUpdateLogic {
	return &SchedulingTypeUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SchedulingTypeUpdateLogic) SchedulingTypeUpdate(in *schedulerclient.SchedulingTypeUpdateReq) (*schedulerclient.CommonResp, error) {
	res, err := l.svcCtx.SchedulingTypeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("没有该排班类型")
		}
		return nil, err
	}

	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	if in.Name != "" {
		res.Name = in.Name
	}

	if len(in.StartTime) > 0 {
		res.StartTime = in.StartTime
	}

	if len(in.EndTime) > 0 {
		res.EndTime = in.EndTime
	}

	if len(in.Colour) > 0 {
		res.Colour.String = in.Colour
		res.Colour.Valid = true
	}

	if in.Remark != "" {
		res.Remark.String = in.Remark
		res.Remark.Valid = true
	}
	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SchedulingTypeModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

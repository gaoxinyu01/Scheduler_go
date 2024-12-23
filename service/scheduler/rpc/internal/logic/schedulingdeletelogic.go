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

type SchedulingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSchedulingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingDeleteLogic {
	return &SchedulingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SchedulingDeleteLogic) SchedulingDelete(in *schedulerclient.SchedulingDeleteReq) (*schedulerclient.CommonResp, error) {
	res, err := l.svcCtx.SchedulingModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("没有该排班列表")
		}
		return nil, err
	}

	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}
	res.DeletedAt.Time = time.Now()
	res.DeletedAt.Valid = true
	res.DeletedName.String = in.DeletedName
	res.DeletedName.Valid = in.DeletedName != ""

	err = l.svcCtx.SchedulingModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

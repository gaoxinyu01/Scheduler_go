package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingTypeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSchedulingTypeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingTypeDeleteLogic {
	return &SchedulingTypeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SchedulingTypeDeleteLogic) SchedulingTypeDelete(in *schedulerclient.SchedulingTypeDeleteReq) (*schedulerclient.CommonResp, error) {
	res, err := l.svcCtx.SchedulingTypeModel.FindOne(l.ctx, in.Id)

	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("没有排班类型")
		}
		return nil, err
	}

	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	err = l.svcCtx.SchedulingTypeModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &schedulerclient.CommonResp{}, nil
}

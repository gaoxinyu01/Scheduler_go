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

type TeamTypeUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTeamTypeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamTypeUpdateLogic {
	return &TeamTypeUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TeamTypeUpdateLogic) TeamTypeUpdate(in *schedulerclient.TeamTypeUpdateReq) (*schedulerclient.CommonResp, error) {
	res, err := l.svcCtx.TeamTypeModel.FindOne(l.ctx, in.Id)

	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("没有该部门")
		}
		return nil, err
	}

	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	if in.Name != "" {
		res.Name = in.Name
	}

	if in.Description != "" {
		res.Description = in.Description
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.TeamTypeModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}
	return &schedulerclient.CommonResp{}, nil
}

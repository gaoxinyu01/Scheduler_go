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

type TeamUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTeamUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamUpdateLogic {
	return &TeamUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TeamUpdateLogic) TeamUpdate(in *schedulerclient.TeamUpdateReq) (*schedulerclient.CommonResp, error) {
	res, err := l.svcCtx.TeamModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("没有该数据")
		}
		return nil, err
	}

	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	if in.UserId != "" {
		res.UserId = in.UserId
	}

	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true
	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true

	err = l.svcCtx.TeamModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

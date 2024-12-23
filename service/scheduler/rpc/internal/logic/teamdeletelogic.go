package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeamDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTeamDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamDeleteLogic {
	return &TeamDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TeamDeleteLogic) TeamDelete(in *schedulerclient.TeamDeleteReq) (*schedulerclient.CommonResp, error) {
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
	// 开启事务
	err = l.svcCtx.TeamModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {

		all, err := l.svcCtx.SchedulingModel.FindListByUserId(res.UserId, in.TenantId, 0)
		if err != nil {
			return err
		}
		for _, item := range *all {
			err = l.svcCtx.SchedulingModel.TransDelete(ctx, sqlx, item.Id)
			if err != nil {
				return err
			}
		}
		err = l.svcCtx.TeamModel.TransDelete(ctx, sqlx, in.Id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

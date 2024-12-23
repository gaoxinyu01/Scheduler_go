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

type TeamTypeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTeamTypeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamTypeDeleteLogic {
	return &TeamTypeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TeamTypeDeleteLogic) TeamTypeDelete(in *schedulerclient.TeamTypeDeleteReq) (*schedulerclient.CommonResp, error) {
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

	// 开启事务
	err = l.svcCtx.TeamTypeModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {

		// 查询team里的用户
		teamUsers, err := l.svcCtx.TeamModel.FindList(0, 0, "", "", "", "", res.Id, in.TenantId)
		if err != nil {
			return err
		}
		for _, teamUser := range *teamUsers {
			// 查询用户未完成的排班任务
			all, err := l.svcCtx.SchedulingModel.FindListByUserId(teamUser.Uid, in.TenantId, 0)
			if err != nil {
				return err
			}
			// 删除用户未完成的排班任务
			for _, item := range *all {
				err = l.svcCtx.SchedulingModel.TransDelete(ctx, sqlx, item.Id)
				if err != nil {
					return err
				}
			}
			// 删除用户
			err = l.svcCtx.TeamModel.TransDelete(ctx, sqlx, teamUser.Id)
			if err != nil {
				return err
			}
		}
		// 删除部门
		err = l.svcCtx.TeamTypeModel.TransDelete(ctx, sqlx, res.Id)
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

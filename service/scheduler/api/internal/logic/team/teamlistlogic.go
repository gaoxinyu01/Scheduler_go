package team

import (
	"Scheduler_go/common"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"

	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeamListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeamListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamListLogic {
	return &TeamListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeamListLogic) TeamList(req *types.TeamListRequest) (resp *types.Response, err error) {
	// 用户登录信息

	all, err := l.svcCtx.SchedulerRpc.TeamFindList(l.ctx, &schedulerclient.TeamFindListReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		NickName:   req.NickName,
		Major:      req.Major,
		Position:   req.Position,
		Telephone:  req.Telephone,
		TeamTypeId: req.TeamTypeId,
	})

	if err != nil {
		return nil, common.NewDefaultError(err.Error())

	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: all,
	}, nil
}

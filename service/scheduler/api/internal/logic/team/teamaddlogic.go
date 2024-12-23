package team

import (
	"Scheduler_go/common"
	"Scheduler_go/common/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/manage/authentication/authenticationclient"
	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeamAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeamAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamAddLogic {
	return &TeamAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeamAddLogic) TeamAdd(req *types.TeamAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	for _, v := range req.UserIds {
		_, err := l.svcCtx.AuthenticationRpc.SysUserFindOne(l.ctx, &authenticationclient.SysUserFindOneReq{
			Id: v,
		})
		if err != nil {
			return nil, common.NewDefaultError(err.Error())
		}
	}

	_, err = l.svcCtx.SchedulerRpc.TeamAdd(l.ctx, &schedulerclient.TeamAddReq{
		CreatedName: tokenData.NickName,
		UserIds:     req.UserIds,
		TeamTypeId:  req.TeamTypeId,
		TenantId:    tokenData.TenantId,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}

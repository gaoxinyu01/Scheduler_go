package team

import (
	"Scheduler_go/common/global/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/manage/authentication/authenticationclient"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"

	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotTeamUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotTeamUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotTeamUserLogic {
	return &NotTeamUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotTeamUserLogic) NotTeamUser(req *types.NotTeamUserRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	// 获取所有用户
	userList, err := l.svcCtx.AuthenticationRpc.SysUserList(l.ctx, &authenticationclient.SysUserListReq{
		ProductId: req.ProductId,
		TenantId:  tokenData.TenantId,
	})

	// 获取所有Team用户
	teamUserList, err := l.svcCtx.SchedulerRpc.TeamFindList(l.ctx, &schedulerclient.TeamFindListReq{
		TenantId: tokenData.TenantId,
	})

	var userListRes []*authenticationclient.SysUserListData
	for _, user := range userList.List {
		state := false
		for _, team := range teamUserList.List {
			if team.UserId == user.Id {
				state = true
				break
			}
		}
		if state == false {
			userListRes = append(userListRes, user)
		}

	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: userListRes,
	}, nil
}

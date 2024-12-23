package scheduling

import (
	"Scheduler_go/common"
	"Scheduler_go/common/global/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/manage/authentication/authenticationclient"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"

	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSchedulingUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingUpLogic {
	return &SchedulingUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SchedulingUpLogic) SchedulingUp(req *types.SchedulingUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	res, err := l.svcCtx.AuthenticationRpc.SysUserFindOne(l.ctx, &authenticationclient.SysUserFindOneReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	_, err = l.svcCtx.SchedulerRpc.SchedulingUpdate(l.ctx, &schedulerclient.SchedulingUpdateReq{
		Id:          req.Id,
		UpdatedName: tokenData.NickName,
		Time:        req.Time,
		Name:        req.Name,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Colour:      req.Colour,
		UserName:    res.NickName,
		TeamId:      req.TeamId,
		UserId:      req.UserId,
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

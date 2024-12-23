package teamType

import (
	"Scheduler_go/common"
	"Scheduler_go/common/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"

	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeamTypeUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeamTypeUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamTypeUpLogic {
	return &TeamTypeUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeamTypeUpLogic) TeamTypeUp(req *types.TeamTypeUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.SchedulerRpc.TeamTypeUpdate(l.ctx, &schedulerclient.TeamTypeUpdateReq{
		Id:          req.Id,
		UpdatedName: tokenData.NickName,
		Name:        req.Name,
		Description: req.Description,
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

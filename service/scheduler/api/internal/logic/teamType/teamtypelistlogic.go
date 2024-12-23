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

type TeamTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeamTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamTypeListLogic {
	return &TeamTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeamTypeListLogic) TeamTypeList(req *types.TeamTypeListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.SchedulerRpc.TeamTypeFindList(l.ctx, &schedulerclient.TeamTypeFindListReq{
		Current:     req.Current,
		PageSize:    req.PageSize,
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
		Data: all,
	}, nil
}

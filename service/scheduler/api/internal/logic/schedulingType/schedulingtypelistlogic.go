package schedulingType

import (
	"Scheduler_go/common"
	"Scheduler_go/common/global/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"

	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSchedulingTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingTypeListLogic {
	return &SchedulingTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SchedulingTypeListLogic) SchedulingTypeList(req *types.SchedulingTypeListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.SchedulerRpc.SchedulingTypeFindList(l.ctx, &schedulerclient.SchedulingTypeFindListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		TenantId: tokenData.TenantId,
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

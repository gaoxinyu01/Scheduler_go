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

type SchedulingTypeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSchedulingTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingTypeAddLogic {
	return &SchedulingTypeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SchedulingTypeAddLogic) SchedulingTypeAdd(req *types.SchedulingTypeAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	_, err = l.svcCtx.SchedulerRpc.SchedulingTypeAdd(l.ctx, &schedulerclient.SchedulingTypeAddReq{
		CreatedName: tokenData.NickName,
		Name:        req.Name,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Remark:      req.Remark,
		Colour:      req.Colour,
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

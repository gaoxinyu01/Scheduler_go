package attendance

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

type AttendanceDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttendanceDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceDelLogic {
	return &AttendanceDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttendanceDelLogic) AttendanceDel(req *types.AttendanceDelRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.SchedulerRpc.AttendanceDelete(l.ctx, &schedulerclient.AttendanceDeleteReq{
		Id:          req.Id,             // 考勤ID
		DeletedName: tokenData.NickName, // 删除人
		TenantId:    tokenData.TenantId, // 租户ID
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

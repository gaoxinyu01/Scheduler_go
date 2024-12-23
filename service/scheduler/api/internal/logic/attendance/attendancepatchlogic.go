package attendance

import (
	"Scheduler_go/common"
	"Scheduler_go/common/global/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"
	"time"

	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendancePatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttendancePatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendancePatchLogic {
	return &AttendancePatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttendancePatchLogic) AttendancePatch(req *types.AttendancePatchRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.SchedulerRpc.AttendancePatch(l.ctx, &schedulerclient.AttendancePatchReq{
		UserId:       req.UserId,                      // 用户ID
		Date:         time.Now().Format("2006-01-02"), // 考勤日期
		SignOffTime:  req.SignOffTime,                 // 签退时间
		SignOffPhoto: req.SignOffPhoto,                // 签退图片
		TenantId:     tokenData.TenantId,              // 租户ID
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

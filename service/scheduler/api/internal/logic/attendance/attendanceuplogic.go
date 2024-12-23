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

type AttendanceUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttendanceUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceUpLogic {
	return &AttendanceUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttendanceUpLogic) AttendanceUp(req *types.AttendanceUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.SchedulerRpc.AttendanceUpdate(l.ctx, &schedulerclient.AttendanceUpdateReq{
		Id:           req.Id,             // 考勤ID
		Name:         req.Name,           // 考勤人
		UserId:       req.UserId,         // 用户ID
		Date:         req.Date,           // 考勤日期
		CheckInTime:  req.CheckInTime,    // 签到时间
		CheckInPhoto: req.CheckInPhoto,   // 签到图片
		SignOffTime:  req.SignOffTime,    // 签退时间
		SignOffPhoto: req.SignOffPhoto,   // 签退图片
		State:        req.State,          // 考勤状态 上班打卡:1,下班打卡:2,打卡正常:3,打卡异常:4
		UpdatedName:  tokenData.NickName, // 更新人
		TenantId:     tokenData.TenantId, // 租户ID
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

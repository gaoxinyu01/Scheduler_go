package logic

import (
	"context"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendanceFindOneDayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendanceFindOneDayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceFindOneDayLogic {
	return &AttendanceFindOneDayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取某天考勤
func (l *AttendanceFindOneDayLogic) AttendanceFindOneDay(in *schedulerclient.AttendanceFindOneDayReq) (*schedulerclient.AttendanceFindOneDayResp, error) {
	res, err := l.svcCtx.AttendanceModel.FindOneDayByDate(l.ctx, in.Date, in.UserId, in.TenantId)
	if err != nil {
		return &schedulerclient.AttendanceFindOneDayResp{
			Date:         in.Date, //考勤日期
			CheckInTime:  0,       //签到时间
			CheckInPhoto: "",      //签到图片
			SignOffTime:  0,       //签退时间
			SignOffPhoto: "",      //签退图片
			State:        3,       //考勤状态 上班打卡:1,打卡正常:2,打卡异常:3
		}, nil

	}

	return &schedulerclient.AttendanceFindOneDayResp{
		Date:         res.Date,              //考勤日期
		CheckInTime:  res.CheckInTime.Int64, //签到时间
		CheckInPhoto: res.CheckInPhoto,      //签到图片
		SignOffTime:  res.SignOffTime.Int64, //签退时间
		SignOffPhoto: res.SignOffPhoto,      //签退图片
		State:        res.State,             //考勤状态 上班打卡:1,打卡正常:2,打卡异常:3
	}, nil
}

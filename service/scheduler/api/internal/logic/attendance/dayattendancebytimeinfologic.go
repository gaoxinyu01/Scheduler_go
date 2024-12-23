package attendance

import (
	"Scheduler_go/common"
	"Scheduler_go/common/global/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"
	"github.com/jinzhu/copier"

	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DayAttendanceByTimeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDayAttendanceByTimeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DayAttendanceByTimeInfoLogic {
	return &DayAttendanceByTimeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DayAttendanceByTimeInfoLogic) DayAttendanceByTimeInfo(req *types.DayAttendanceByTimeInfoRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	// 获取考勤
	taskCount, err := l.svcCtx.SchedulerRpc.AttendanceByDays(l.ctx, &schedulerclient.AttendanceByDaysReq{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		UserId:    req.UserId,
		TenantId:  tokenData.TenantId,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	var result DayAttendanceByTimeData
	_ = copier.Copy(&result.TaskMonthsCounts, taskCount.List)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type DayAttendanceByTimeData struct {
	TaskMonthsCounts []DayAttendanceByTime `json:"task_attendance_counts"`
}

type DayAttendanceByTime struct {
	Date         string `json:"date"`           // 考勤日期,
	CheckInTime  int64  `json:"check_in_time"`  // 签到时间,
	CheckInPhoto string `json:"check_in_photo"` // 签到图片,
	SignOffTime  int64  `json:"sign_off_time"`  // 签退时间,
	SignOffPhoto string `json:"sign_off_photo"` // 签退图片,
	State        int64  `json:"state"`          // 考勤状态 上班打卡:1,打卡正常:2,打卡异常:3,
}

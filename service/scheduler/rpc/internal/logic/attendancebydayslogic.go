package logic

import (
	"Scheduler_go/common"
	"Scheduler_go/common/global"
	"context"
	"fmt"
	"time"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendanceByDaysLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendanceByDaysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceByDaysLogic {
	return &AttendanceByDaysLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据时间段获取每日考勤
func (l *AttendanceByDaysLogic) AttendanceByDays(in *schedulerclient.AttendanceByDaysReq) (*schedulerclient.AttendanceByDaysResp, error) {
	rangeTime := in.EndTime - in.StartTime

	//考虑到查询单日，这点代码暂时注释
	if rangeTime <= 0 {
		return nil, fmt.Errorf("开始天数不能小于等于结束天数")
	}

	startTime := time.UnixMilli(in.StartTime)

	// 计算差多少天
	days := int(rangeTime / (time.Hour * 24).Milliseconds())
	if days > 367 {
		return nil, common.NewDefaultError("计算日不能超过一年")
	}

	var AttendanceByDays []*schedulerclient.AttendanceByDaysCounts

	for i := 0; i < days; i++ {
		Date := time.Date(startTime.Year(), startTime.Month(), startTime.Day()+i, 0, 0, 0, 0, global.ShangHaiTime).UnixMilli()
		dataReq := time.UnixMilli(Date).Format("2006-01-02")
		res, _ := l.svcCtx.AttendanceModel.FindOneDayByDate(l.ctx, dataReq, in.UserId, in.TenantId)
		if res == nil {
			continue
		}

		AttendanceByDays = append(AttendanceByDays, &schedulerclient.AttendanceByDaysCounts{
			Date:         res.Date,              //考勤日期
			CheckInTime:  res.CheckInTime.Int64, //签到时间
			CheckInPhoto: res.CheckInPhoto,      //签到图片
			SignOffTime:  res.SignOffTime.Int64, //签退时间
			SignOffPhoto: res.SignOffPhoto,      //签退图片
			State:        res.State,             //考勤状态 上班打卡:1,打卡正常:2,打卡异常:3

		})
	}

	return &schedulerclient.AttendanceByDaysResp{
		List: AttendanceByDays,
	}, nil
}

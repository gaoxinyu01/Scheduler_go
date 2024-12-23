package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendanceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendanceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceUpdateLogic {
	return &AttendanceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AttendanceUpdateLogic) AttendanceUpdate(in *schedulerclient.AttendanceUpdateReq) (*schedulerclient.CommonResp, error) {
	res, err := l.svcCtx.AttendanceModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("Attendance没有该ID：" + in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, errors.New("Attendance该ID已被删除：" + in.Id)
	}
	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	// 考勤人
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 用户ID
	if len(in.UserId) > 0 {
		res.UserId = in.UserId
	}
	// 考勤日期
	if len(in.Date) > 0 {
		res.Date = in.Date
	}
	// 签到时间
	if in.CheckInTime != 0 {
		res.CheckInTime.Int64 = in.CheckInTime
		res.CheckInTime.Valid = true
	}
	// 签到图片
	if len(in.CheckInPhoto) > 0 {
		res.CheckInPhoto = in.CheckInPhoto
	}
	// 签退时间
	if in.SignOffTime != 0 {
		res.SignOffTime.Int64 = in.SignOffTime
		res.SignOffTime.Valid = true
	}
	// 签退图片
	if len(in.SignOffPhoto) > 0 {
		res.SignOffPhoto = in.SignOffPhoto
	}
	// 考勤状态 上班打卡:1,打卡正常:2,打卡异常:3
	if in.State != 0 {
		res.State = in.State
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.AttendanceModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &schedulerclient.CommonResp{}, nil
}

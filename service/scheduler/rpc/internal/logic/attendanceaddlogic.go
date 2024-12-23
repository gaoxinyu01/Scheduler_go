package logic

import (
	"Scheduler_go/service/scheduler/model"
	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"
	"database/sql"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendanceAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendanceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceAddLogic {
	return &AttendanceAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 考勤
func (l *AttendanceAddLogic) AttendanceAdd(in *schedulerclient.AttendanceAddReq) (resp *schedulerclient.CommonResp, err error) {
	// 根据用户ID，日期，租户查询签到记录，如果已经签到，不允许签到
	_, err = l.svcCtx.AttendanceModel.FindOneByUserIdAndDate(l.ctx, in.Date, in.UserId, in.TenantId)
	if err == nil {
		return nil, errors.New("已经签到！")
	}

	_, err = l.svcCtx.AttendanceModel.Insert(l.ctx, &model.Attendance{
		Id:           uuid.NewV4().String(),                             // ID
		CreatedAt:    time.Now(),                                        // 创建时间
		Name:         in.Name,                                           // 考勤人
		UserId:       in.UserId,                                         // 用户ID
		Date:         in.Date,                                           // 考勤日期
		CheckInTime:  sql.NullInt64{Int64: in.CheckInTime, Valid: true}, // 签到时间
		CheckInPhoto: in.CheckInPhoto,                                   // 签到图片
		SignOffTime:  sql.NullInt64{Int64: in.SignOffTime, Valid: true}, // 签退时间
		SignOffPhoto: in.SignOffPhoto,                                   // 签退图片
		State:        in.State,                                          // 考勤状态 上班打卡:1,打卡正常:2,打卡异常:3
		CreatedName:  in.CreatedName,                                    // 创建人
		TenantId:     in.TenantId,                                       // 租户ID
	})
	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

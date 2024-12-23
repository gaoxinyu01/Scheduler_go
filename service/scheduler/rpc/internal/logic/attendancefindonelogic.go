package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendanceFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendanceFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceFindOneLogic {
	return &AttendanceFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AttendanceFindOneLogic) AttendanceFindOne(in *schedulerclient.AttendanceFindOneReq) (*schedulerclient.AttendanceFindOneResp, error) {
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

	return &schedulerclient.AttendanceFindOneResp{
		Id:           res.Id,                         //考勤ID
		Name:         res.Name,                       //考勤人
		UserId:       res.UserId,                     //用户ID
		Date:         res.Date,                       //考勤日期
		CheckInTime:  res.CheckInTime.Int64,          //签到时间
		CheckInPhoto: res.CheckInPhoto,               //签到图片
		SignOffTime:  res.SignOffTime.Int64,          //签退时间
		SignOffPhoto: res.SignOffPhoto,               //签退图片
		State:        res.State,                      //考勤状态 上班打卡:1,打卡正常:2,打卡异常:3
		CreatedAt:    res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:    res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName:  res.CreatedName,                //创建人
		UpdatedName:  res.UpdatedName.String,         //更新人
	}, nil
}

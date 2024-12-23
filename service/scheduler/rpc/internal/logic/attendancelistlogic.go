package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendanceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendanceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceListLogic {
	return &AttendanceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AttendanceListLogic) AttendanceList(in *schedulerclient.AttendanceListReq) (*schedulerclient.AttendanceListResp, error) {
	whereBuilder := l.svcCtx.AttendanceModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 考勤人
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 用户ID
	if len(in.UserId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"user_id ": "%" + in.UserId + "%",
		})
	}
	// 考勤日期
	if len(in.Date) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"date ": "%" + in.Date + "%",
		})
	}
	// 签到时间
	if in.CheckInTime != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"check_in_time ": in.CheckInTime,
		})
	}
	// 签到图片
	if len(in.CheckInPhoto) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"check_in_photo ": "%" + in.CheckInPhoto + "%",
		})
	}
	// 签退时间
	if in.SignOffTime != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"sign_off_time ": in.SignOffTime,
		})
	}
	// 签退图片
	if len(in.SignOffPhoto) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"sign_off_photo ": "%" + in.SignOffPhoto + "%",
		})
	}
	// 考勤状态 上班打卡:1,下班打卡:2,打卡正常:3,打卡异常:4
	if in.State != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}

	all, err := l.svcCtx.AttendanceModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.AttendanceModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	countBuilder = countBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 考勤人
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 用户ID
	if len(in.UserId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"user_id ": "%" + in.UserId + "%",
		})
	}
	// 考勤日期
	if len(in.Date) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"date ": "%" + in.Date + "%",
		})
	}
	// 签到时间
	if in.CheckInTime != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"check_in_time ": in.CheckInTime,
		})
	}
	// 签退时间
	if in.SignOffTime != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"sign_off_time ": in.SignOffTime,
		})
	}
	// 考勤状态 上班打卡:1,下班打卡:2,打卡正常:3,打卡异常:4
	if in.State != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}
	count, err := l.svcCtx.AttendanceModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*schedulerclient.AttendanceListData
	for _, item := range all {
		list = append(list, &schedulerclient.AttendanceListData{
			Id:           item.Id,                         //考勤ID
			Name:         item.Name,                       //考勤人
			UserId:       item.UserId,                     //用户ID
			Date:         item.Date,                       //考勤日期
			CheckInTime:  item.CheckInTime.Int64,          //签到时间
			CheckInPhoto: item.CheckInPhoto,               //签到图片
			SignOffTime:  item.SignOffTime.Int64,          //签退时间
			SignOffPhoto: item.SignOffPhoto,               //签退图片
			State:        item.State,                      //考勤状态 上班打卡:1,下班打卡:2,打卡正常:3,打卡异常:4
			CreatedAt:    item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:    item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:  item.CreatedName,                //创建人
			UpdatedName:  item.UpdatedName.String,         //更新人
		})
	}

	return &schedulerclient.AttendanceListResp{
		Total: count,
		List:  list,
	}, nil
}

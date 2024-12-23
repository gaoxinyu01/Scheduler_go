package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendancePatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendancePatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendancePatchLogic {
	return &AttendancePatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 签退
func (l *AttendancePatchLogic) AttendancePatch(in *schedulerclient.AttendancePatchReq) (*schedulerclient.CommonResp, error) {
	Date := in.Date
	if l.svcCtx.Config.IsOvernight {
		// 调用函数并打印结果
		Date, _ = CalculateYesterday(in.Date)

	}

	// 根据用户ID，日期，租户查询签到记录
	res, err := l.svcCtx.AttendanceModel.FindOneByUserIdAndDate(l.ctx, Date, in.UserId, in.TenantId)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("未签到！")
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, errors.New("Attendance该ID已被删除：" + in.UserId)
	}
	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
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

	timeDiff := in.SignOffTime - res.CheckInTime.Int64 - l.svcCtx.Config.WorkingSystem*3600000
	if timeDiff < 0 {
		res.State = 3

	} else {
		res.State = 2
	}

	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.AttendanceModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &schedulerclient.CommonResp{}, nil
}

// 根据字符串格式的日历日期计算并返回昨日日期的字符串
func CalculateYesterday(inputDate string) (string, error) {
	// 定义日期格式
	const layout = "2006-01-02"

	// 解析输入字符串为time.Time类型
	parsedTime, err := time.Parse(layout, inputDate)
	if err != nil {
		return "", fmt.Errorf("failed to parse date: %w", err)
	}

	// 计算昨日的时间
	yesterday := parsedTime.Add(-24 * time.Hour)

	// 将结果格式化为字符串
	yesterdayStr := yesterday.Format(layout)

	return yesterdayStr, nil
}

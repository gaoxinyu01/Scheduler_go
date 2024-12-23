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

type AttendanceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttendanceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceAddLogic {
	return &AttendanceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttendanceAddLogic) AttendanceAdd(req *types.AttendanceAddRequest) (resp *types.Response, err error) {
	// 用户登录信息

	tokenData := jwtx.ParseToken(l.ctx)
	_, err = l.svcCtx.SchedulerRpc.AttendanceAdd(l.ctx, &schedulerclient.AttendanceAddReq{
		Name:   tokenData.NickName, // 考勤人
		UserId: req.UserId,         // 用户ID
		//Date:         time.Now().Format("2006-01-02"), // 考勤日期   这种方式更严谨
		Date:         time.UnixMilli(req.CheckInTime).Format("2006-01-02"), // 便于开发测试，这里改为获取签到时间的日期
		CheckInTime:  req.CheckInTime,                                      // 签到时间
		CheckInPhoto: req.CheckInPhoto,                                     // 签到图片
		State:        1,                                                    //签到打卡
		CreatedName:  tokenData.NickName,                                   // 创建人
		TenantId:     tokenData.TenantId,                                   // 租户ID
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

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

type AttendanceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttendanceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceInfoLogic {
	return &AttendanceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttendanceInfoLogic) AttendanceInfo(req *types.AttendanceInfoRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	res, err := l.svcCtx.SchedulerRpc.AttendanceFindOne(l.ctx, &schedulerclient.AttendanceFindOneReq{
		Id:       req.Id,             // 考勤ID
		TenantId: tokenData.TenantId, // 租户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result AttendanceFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type AttendanceFindOneResp struct {
	Id           string `json:"id"`             // 考勤ID,
	Name         string `json:"name"`           // 考勤人,
	UserId       string `json:"user_id"`        // 用户ID,
	Date         string `json:"date"`           // 考勤日期,
	CheckInTime  int64  `json:"check_in_time"`  // 签到时间,
	CheckInPhoto string `json:"check_in_photo"` // 签到图片,
	SignOffTime  int64  `json:"sign_off_time"`  // 签退时间,
	SignOffPhoto string `json:"sign_off_photo"` // 签退图片,
	State        int64  `json:"state"`          // 考勤状态 上班打卡:1,打卡正常:2,打卡异常:3,
	CreatedAt    int64  `json:"created_at"`     // 创建时间,
	UpdatedAt    int64  `json:"updated_at"`     // 更新时间,
	CreatedName  string `json:"created_name"`   // 创建人,
	UpdatedName  string `json:"updated_name"`   // 更新人
}

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

type AttendanceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttendanceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceListLogic {
	return &AttendanceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttendanceListLogic) AttendanceList(req *types.AttendanceListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.SchedulerRpc.AttendanceList(l.ctx, &schedulerclient.AttendanceListReq{
		Current:     req.Current,        // 页码
		PageSize:    req.PageSize,       // 页数
		Name:        req.Name,           // 考勤人
		UserId:      req.UserId,         // 用户ID
		Date:        req.Date,           // 考勤日期
		CheckInTime: req.CheckInTime,    // 签到时间
		SignOffTime: req.SignOffTime,    // 签退时间
		State:       req.State,          // 考勤状态 上班打卡:1,下班打卡:2,打卡正常:3,打卡异常:4
		TenantId:    tokenData.TenantId, // 租户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result AttendanceListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type AttendanceListResp struct {
	Total int64                 `json:"total"`
	List  []*AttendanceDataList `json:"list"`
}

type AttendanceDataList struct {
	Name         string `json:"name"`           // 考勤人,
	Date         string `json:"date"`           // 考勤日期,
	CheckInTime  int64  `json:"check_in_time"`  // 签到时间,
	CheckInPhoto string `json:"check_in_photo"` // 签到图片,check_in_photo
	SignOffTime  int64  `json:"sign_off_time"`  // 签退时间,
	SignOffPhoto string `json:"sign_off_photo"` // 签退图片,
	State        int64  `json:"state"`          // 考勤状态 上班打卡:1,打卡正常:2,打卡异常:3,
}

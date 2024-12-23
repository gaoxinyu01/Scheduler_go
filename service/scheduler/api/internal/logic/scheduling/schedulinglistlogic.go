package scheduling

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

type SchedulingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSchedulingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingListLogic {
	return &SchedulingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SchedulingListLogic) SchedulingList(req *types.SchedulingListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.SchedulerRpc.SchedulingFindList(l.ctx, &schedulerclient.SchedulingFindListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		Time:      req.Time,
		Name:      req.Name,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		TeamName:  req.TeamName,
		UserName:  req.UserName,
		TenantId:  tokenData.TenantId,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SchedulingFindListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

// 查询排班列表返回
type SchedulingFindListResp struct {
	Total int64                     `json:"total"`
	List  []*SchedulingFindListData `json:"list"`
}

type SchedulingFindListData struct {
	Id           string `json:"id"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	CreatedName  string `json:"created_name"`
	UpdatedName  string `json:"updated_name"`
	Time         int64  `json:"time"`       //  排班时间
	Name         string `json:"name"`       //  用户ID
	StartTime    int64  `json:"start_time"` // 部门ID
	EndTime      int64  `json:"end_time"`   // 租户ID
	Colour       string `json:"colour"`     // 租户ID
	TeamName     string `json:"team_name"`  //备注
	UserName     string `json:"user_name"`
	JobStartTime int64  `json:"job_start_time"`
	JobEndTime   int64  `json:"job_end_time"`
}

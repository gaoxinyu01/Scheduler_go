package procTask

import (
	"Scheduler_go/common"
	"Scheduler_go/common/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/workflow/rpc/workflowclient"
	"context"
	"github.com/jinzhu/copier"

	"Scheduler_go/service/workflow/api/internal/svc"
	"Scheduler_go/service/workflow/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskFinishedListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskFinishedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskFinishedListLogic {
	return &TaskFinishedListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskFinishedListLogic) TaskFinishedList(req *types.TaskFinishedListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.WorkflowRpc.TaskFinishedList(l.ctx, &workflowclient.TaskFinishedListReq{
		Current:         req.Current,        // 页码
		PageSize:        req.PageSize,       // 页数
		UserId:          req.UserId,         // 分配用户ID
		ProceName:       req.ProcName,       //
		TenantId:        tokenData.TenantId, // 租户ID
		Data:            req.Data,           //
		CreatedName:     tokenData.NickName,
		SortByAsc:       req.SortByAsc,
		IgnoreStartByMe: req.IgnoreStartByMe,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TaskFinishedListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TaskFinishedListResp struct {
	Total int64               `json:"total"`
	List  []*ProcTaskDataList `json:"list"`
}

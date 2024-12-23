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

type TaskUpstreamNodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskUpstreamNodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskUpstreamNodeListLogic {
	return &TaskUpstreamNodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskUpstreamNodeListLogic) TaskUpstreamNodeList(req *types.TaskUpstreamNodeListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.WorkflowRpc.TaskUpstreamNodeList(l.ctx, &workflowclient.TaskUpstreamNodeListReq{
		Taskid:   req.TaskId,
		TenantId: tokenData.TenantId, // 租户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TaskUpstreamNodeListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TaskUpstreamNodeListResp struct {
	Total int64               `json:"total"`
	List  []*ProcTaskDataList `json:"list"`
}

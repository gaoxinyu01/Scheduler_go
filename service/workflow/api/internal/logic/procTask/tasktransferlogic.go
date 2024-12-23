package procTask

import (
	"Scheduler_go/common"
	"Scheduler_go/common/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/workflow/rpc/workflowclient"
	"context"

	"Scheduler_go/service/workflow/api/internal/svc"
	"Scheduler_go/service/workflow/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskTransferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskTransferLogic {
	return &TaskTransferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskTransferLogic) TaskTransfer(req *types.TaskTransferRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.WorkflowRpc.TaskTransfer(l.ctx, &workflowclient.TaskTransferReq{
		TaskId:      req.TaskId,
		Users:       req.Users,
		TenantId:    tokenData.TenantId, // 租户ID
		Data:        req.Data,           //
		CreatedName: tokenData.NickName, // 创建人
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
package procInst

import (
	"Scheduler_go/common"
	"Scheduler_go/common/global/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/workflow/rpc/workflowclient"
	"context"

	"Scheduler_go/service/workflow/api/internal/svc"
	"Scheduler_go/service/workflow/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcInstUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcInstUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstUpLogic {
	return &ProcInstUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcInstUpLogic) ProcInstUp(req *types.ProcInstUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.WorkflowRpc.ProcInstUpdate(l.ctx, &workflowclient.ProcInstUpdateReq{
		Id:            req.Id,             // 流程实例ID
		ProcId:        req.ProcId,         // 流程ID
		ProcName:      req.ProcName,       // 流程名称
		ProcVersion:   req.ProcVersion,    // 流程版本号
		BusinessId:    req.BusinessId,     // 业务ID
		Starter:       req.Starter,        // 流程发起人用户ID
		CurrentNodeId: req.CurrentNodeId,  // 当前进行节点ID
		VariablesJson: req.VariablesJson,  // 变量(Json)
		Status:        req.Status,         // 状态 0 未完成（审批中） 1 已完成 2 撤销
		TenantId:      tokenData.TenantId, // 租户ID
		Data:          req.Data,           //
		UpdatedName:   tokenData.NickName, // 更新人
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

package procInst

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

type ProcInstRevokeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcInstRevokeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstRevokeLogic {
	return &ProcInstRevokeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcInstRevokeLogic) ProcInstRevoke(req *types.ProcInstRevokeRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)
	res, err := l.svcCtx.WorkflowRpc.ProcInstFindOne(l.ctx, &workflowclient.ProcInstFindOneReq{
		Id:       req.Id,             // 流程实例ID
		TenantId: tokenData.TenantId, // 租户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	_, err = l.svcCtx.WorkflowRpc.ProcInstRevoke(l.ctx, &workflowclient.ProcInstRevokeReq{
		Id:           req.Id, // 流程实例ID
		RevokeUserId: req.RevokeUserID,
		Force:        req.Force,
		TenantId:     tokenData.TenantId, // 租户ID
		Data:         res.Data,           //
		UpdatedName:  tokenData.NickName, // 更新人
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

package procInst

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

type ProcInstTaskHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcInstTaskHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstTaskHistoryLogic {
	return &ProcInstTaskHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcInstTaskHistoryLogic) ProcInstTaskHistory(req *types.ProcInstTaskHistoryRequest) (resp *types.Response, err error) {

	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	res, err := l.svcCtx.WorkflowRpc.procInstTaskHistory(l.ctx, &workflowclient.procInstTaskHistoryReq{
		Id:       req.InstId,         // 流程实例ID
		TenantId: tokenData.TenantId, // 租户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result PprocInstTaskHistoryResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

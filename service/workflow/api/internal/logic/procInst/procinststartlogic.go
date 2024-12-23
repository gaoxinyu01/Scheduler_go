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

type ProcInstStartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcInstStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstStartLogic {
	return &ProcInstStartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcInstStartLogic) ProcInstStart(req *types.ProcInstStartRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.WorkflowRpc.ProcInstStart(l.ctx, &workflowclient.ProcInstStartReq{
		ProcId:        req.ProcId,         // 流程ID
		BusinessId:    req.BusinessId,     // 业务ID
		VariablesJson: req.VariablesJson,  // 变量(Json)
		Comment:       req.Comment,        //评论意见
		TenantId:      tokenData.TenantId, // 租户ID
		Data:          req.Data,           //
		CreatedName:   tokenData.NickName, // 创建人
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

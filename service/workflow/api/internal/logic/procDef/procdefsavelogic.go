package procDef

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

type ProcDefSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcDefSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcDefSaveLogic {
	return &ProcDefSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcDefSaveLogic) ProcDefSave(req *types.ProcDefSaveRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.WorkflowRpc.ProcDefSave(l.ctx, &workflowclient.ProcDefSaveReq{
		Resource:     req.Resource, // 流程定义模板
		CreateUserId: req.CreateUserId,
		TenantId:     tokenData.TenantId, // 租户ID
		CreatedName:  tokenData.NickName, // 创建人
		Data:         req.Data,           //预留
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

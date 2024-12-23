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

type ProcDefUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcDefUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcDefUpLogic {
	return &ProcDefUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcDefUpLogic) ProcDefUp(req *types.ProcDefUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.WorkflowRpc.ProcDefUpdate(l.ctx, &workflowclient.ProcDefUpdateReq{
		Id:           req.Id,             // 流程模板ID
		Name:         req.Name,           // 流程名称
		Version:      req.Version,        // 版本号
		ProcType:     req.ProcType,       // 流程类型
		Resource:     req.Resource,       // 流程定义模板
		CreateUserId: req.CreateUserId,   // 创建者ID
		Source:       req.Source,         // 来源
		TenantId:     tokenData.TenantId, // 租户ID
		Data:         req.Data,           //
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

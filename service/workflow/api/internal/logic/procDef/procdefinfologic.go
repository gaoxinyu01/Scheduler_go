package procDef

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

type ProcDefInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcDefInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcDefInfoLogic {
	return &ProcDefInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcDefInfoLogic) ProcDefInfo(req *types.ProcDefInfoRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	res, err := l.svcCtx.WorkflowRpc.ProcDefFindOne(l.ctx, &workflowclient.ProcDefFindOneReq{
		Id:       req.Id,             // 流程模板ID
		TenantId: tokenData.TenantId, // 租户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result ProcDefFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type ProcDefFindOneResp struct {
	Id           int64  `json:"id"`             // 流程模板ID,
	Name         string `json:"name"`           // 流程名称,
	Version      int64  `json:"version"`        // 版本号,
	ProcType     int64  `json:"proc_type"`      // 流程类型,
	Resource     string `json:"resource"`       // 流程定义模板,
	CreateUserId string `json:"create_user_id"` // 创建者ID,
	Source       string `json:"source"`         // 来源,
	Data         string `json:"data"`           // ,
	CreatedAt    int64  `json:"created_at"`     // 创建时间,
	UpdatedAt    int64  `json:"updated_at"`     // 更新时间,
	CreatedName  string `json:"created_name"`   // 创建人,
	UpdatedName  string `json:"updated_name"`   // 更新人
}

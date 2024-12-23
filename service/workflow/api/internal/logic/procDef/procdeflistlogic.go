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

type ProcDefListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcDefListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcDefListLogic {
	return &ProcDefListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcDefListLogic) ProcDefList(req *types.ProcDefListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.WorkflowRpc.ProcDefList(l.ctx, &workflowclient.ProcDefListReq{
		Current:      req.Current,        // 页码
		PageSize:     req.PageSize,       // 页数
		Name:         req.Name,           // 流程名称
		Version:      req.Version,        // 版本号
		ProcType:     req.ProcType,       // 流程类型
		Resource:     req.Resource,       // 流程定义模板
		CreateUserId: req.CreateUserId,   // 创建者ID
		Source:       req.Source,         // 来源
		TenantId:     tokenData.TenantId, // 租户ID
		Data:         req.Data,           //
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result ProcDefListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type ProcDefListResp struct {
	Total int64              `json:"total"`
	List  []*ProcDefDataList `json:"list"`
}

type ProcDefDataList struct {
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

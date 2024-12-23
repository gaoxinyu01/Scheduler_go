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

type ProcInstListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcInstListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstListLogic {
	return &ProcInstListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcInstListLogic) ProcInstList(req *types.ProcInstListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.WorkflowRpc.ProcInstList(l.ctx, &workflowclient.ProcInstListReq{
		Current:       req.Current,        // 页码
		PageSize:      req.PageSize,       // 页数
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
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result ProcInstListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type ProcInstListResp struct {
	Total int64               `json:"total"`
	List  []*ProcInstDataList `json:"list"`
}

type ProcInstDataList struct {
	Id            int64  `json:"id"`              // 流程实例ID,
	ProcId        int64  `json:"proc_id"`         // 流程ID,
	ProcName      string `json:"proc_name"`       // 流程名称,
	ProcVersion   int64  `json:"proc_version"`    // 流程版本号,
	BusinessId    string `json:"business_id"`     // 业务ID,
	Starter       string `json:"starter"`         // 流程发起人用户ID,
	CurrentNodeId string `json:"current_node_id"` // 当前进行节点ID,
	VariablesJson string `json:"variables_json"`  // 变量(Json),
	Status        int64  `json:"status"`          // 状态 0 未完成（审批中） 1 已完成 2 撤销,
	Data          string `json:"data"`            // ,
	CreatedAt     int64  `json:"created_at"`      // 创建时间,
	UpdatedAt     int64  `json:"updated_at"`      // 更新时间,
	CreatedName   string `json:"created_name"`    // 创建人,
	UpdatedName   string `json:"updated_name"`    // 更新人
}

package procTask

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

type TaskToDoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskToDoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskToDoListLogic {
	return &TaskToDoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskToDoListLogic) TaskToDoList(req *types.TaskToDoListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.WorkflowRpc.TaskToDoList(l.ctx, &workflowclient.TaskToDoListReq{
		Current:     req.Current,        // 页码
		PageSize:    req.PageSize,       // 页数
		UserId:      req.UserId,         // 分配用户ID
		ProceId:     req.ProcId,         //
		TenantId:    tokenData.TenantId, // 租户ID
		Data:        req.Data,           //
		CreatedName: tokenData.NickName,
		SortByAsc:   req.SortByAsc,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TaskToDoListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TaskToDoListResp struct {
	Total int64               `json:"total"`
	List  []*ProcTaskDataList `json:"list"`
}

type ProcTaskDataList struct {
	Id                 int64  `json:"id"`                    // 任务ID,
	ProcId             int64  `json:"proc_id"`               // 流程ID,
	ProcInstId         int64  `json:"proc_inst_id"`          // 流程实例ID,
	BusinessId         string `json:"business_id"`           // 业务ID,
	Starter            string `json:"starter"`               // 流程发起人用户ID,
	NodeId             string `json:"node_id"`               // 节点ID,
	NodeName           string `json:"node_name"`             // 节点名称,
	PrevNodeId         string `json:"prev_node_id"`          // 上个处理节点ID,
	IsCosigned         int64  `json:"is_cosigned"`           // 任意一人通过即可 1:会签,
	BatchCode          string `json:"batch_code"`            // 批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别\",
	UserId             string `json:"user_id"`               // 分配用户ID,
	Status             int64  `json:"status"`                // 任务状态:0:初始 1:通过 2:驳回,
	IsFinished         int64  `json:"is_finished"`           // 0:任务未完成 1:处理完成,
	Comment            string `json:"comment"`               // 任务备注,
	ProcInstCreateTime int64  `json:"proc_inst_create_time"` // 流程实例创建时间,
	FinishedTime       int64  `json:"finished_time"`         // 处理任务时间,
	Data               string `json:"data"`                  // ,
	CreatedAt          int64  `json:"created_at"`            // 创建时间,
	UpdatedAt          int64  `json:"updated_at"`            // 更新时间,
	CreatedName        string `json:"created_name"`          // 创建人,
	UpdatedName        string `json:"updated_name"`          // 更新人
}

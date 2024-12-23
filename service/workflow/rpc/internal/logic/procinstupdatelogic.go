package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcInstUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcInstUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstUpdateLogic {
	return &ProcInstUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcInstUpdateLogic) ProcInstUpdate(in *workflowclient.ProcInstUpdateReq) (resp *workflowclient.CommonResp, err error) {

	res, err := l.svcCtx.ProcInstModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcInst没有该ID：%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("ProcInst该ID已被删除：%v", in.Id)
	}
	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	// 流程ID
	if in.ProcId != 0 {
		res.ProcId = in.ProcId
	}
	// 流程名称
	if len(in.ProcName) > 0 {
		res.ProcName = in.ProcName
	}
	// 流程版本号
	if in.ProcVersion != 0 {
		res.ProcVersion = in.ProcVersion
	}
	// 业务ID
	if len(in.BusinessId) > 0 {
		res.BusinessId = in.BusinessId
	}
	// 流程发起人用户ID
	if len(in.Starter) > 0 {
		res.Starter = in.Starter
	}
	// 当前进行节点ID
	if len(in.CurrentNodeId) > 0 {
		res.CurrentNodeId = in.CurrentNodeId
	}
	// 变量(Json)
	if len(in.VariablesJson) > 0 {
		res.VariablesJson.String = in.VariablesJson
		res.VariablesJson.Valid = true
	}
	// 状态 0 未完成（审批中） 1 已完成 2 撤销
	if in.Status != 0 {
		res.Status = in.Status
	}
	//
	if len(in.Data) > 0 {
		res.Data.String = in.Data
		res.Data.Valid = true
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.ProcInstModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &workflowclient.CommonResp{}, nil

}

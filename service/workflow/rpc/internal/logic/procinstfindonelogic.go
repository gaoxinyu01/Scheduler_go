package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcInstFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcInstFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstFindOneLogic {
	return &ProcInstFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcInstFindOneLogic) ProcInstFindOne(in *workflowclient.ProcInstFindOneReq) (resp *workflowclient.ProcInstFindOneResp, err error) {

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

	return &workflowclient.ProcInstFindOneResp{
		Id:            res.Id,                         //流程实例ID
		ProcId:        res.ProcId,                     //流程ID
		ProcName:      res.ProcName,                   //流程名称
		ProcVersion:   res.ProcVersion,                //流程版本号
		BusinessId:    res.BusinessId,                 //业务ID
		Starter:       res.Starter,                    //流程发起人用户ID
		CurrentNodeId: res.CurrentNodeId,              //当前进行节点ID
		VariablesJson: res.VariablesJson.String,       //变量(Json)
		Status:        res.Status,                     //状态 0 未完成（审批中） 1 已完成 2 撤销
		Data:          res.Data.String,                //
		CreatedAt:     res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:     res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName:   res.CreatedName,                //创建人
		UpdatedName:   res.UpdatedName.String,         //更新人
	}, nil

}

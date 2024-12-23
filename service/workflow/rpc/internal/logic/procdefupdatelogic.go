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

type ProcDefUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcDefUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcDefUpdateLogic {
	return &ProcDefUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcDefUpdateLogic) ProcDefUpdate(in *workflowclient.ProcDefUpdateReq) (resp *workflowclient.CommonResp, err error) {

	res, err := l.svcCtx.ProcDefModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcDef没有该ID：%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("ProcDef该ID已被删除：%v", in.Id)
	}
	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	// 流程名称
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 版本号
	if in.Version != 0 {
		res.Version = in.Version
	}
	// 流程类型
	if in.ProcType != 0 {
		res.ProcType = in.ProcType
	}
	// 流程定义模板
	if len(in.Resource) > 0 {
		res.Resource = in.Resource
	}
	// 创建者ID
	if len(in.CreateUserId) > 0 {
		res.CreateUserId.String = in.CreateUserId
		res.CreateUserId.Valid = true
	}
	// 来源
	if len(in.Source) > 0 {
		res.Source.String = in.Source
		res.Source.Valid = true
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

	err = l.svcCtx.ProcDefModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &workflowclient.CommonResp{}, nil

}

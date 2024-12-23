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

type ProcDefFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcDefFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcDefFindOneLogic {
	return &ProcDefFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcDefFindOneLogic) ProcDefFindOne(in *workflowclient.ProcDefFindOneReq) (resp *workflowclient.ProcDefFindOneResp, err error) {

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

	return &workflowclient.ProcDefFindOneResp{
		Id:           res.Id,                         //流程模板ID
		Name:         res.Name,                       //流程名称
		Version:      res.Version,                    //版本号
		ProcType:     res.ProcType,                   //流程类型
		Resource:     res.Resource,                   //流程定义模板
		CreateUserId: res.CreateUserId.String,        //创建者ID
		Source:       res.Source.String,              //来源
		Data:         res.Data.String,                //
		CreatedAt:    res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:    res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName:  res.CreatedName,                //创建人
		UpdatedName:  res.UpdatedName.String,         //更新人
	}, nil

}

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

type ProcInstDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcInstDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstDeleteLogic {
	return &ProcInstDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcInstDeleteLogic) ProcInstDelete(in *workflowclient.ProcInstDeleteReq) (resp *workflowclient.CommonResp, err error) {

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

	res.DeletedAt.Time = time.Now()
	res.DeletedAt.Valid = true
	res.DeletedName.String = in.DeletedName
	res.DeletedName.Valid = true

	err = l.svcCtx.ProcInstModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}

	return &workflowclient.CommonResp{}, nil
}

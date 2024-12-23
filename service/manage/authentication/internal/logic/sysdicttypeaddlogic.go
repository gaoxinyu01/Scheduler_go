package logic

import (
	"Scheduler_go/service/manage/authentication/authenticationclient"
	"Scheduler_go/service/manage/authentication/internal/svc"
	"Scheduler_go/service/manage/authentication/model"
	"context"
	"database/sql"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictTypeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeAddLogic {
	return &SysDictTypeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 字典类型
func (l *SysDictTypeAddLogic) SysDictTypeAdd(in *authenticationclient.SysDictTypeAddReq) (*authenticationclient.CommonResp, error) {

	_, err := l.svcCtx.SysDictTypeModel.Insert(l.ctx, &model.SysDictType{
		CreatedAt:   time.Now(),                                                // 创建时间
		CreatedName: in.CreatedName,                                            // 创建人
		Name:        in.Name,                                                   // 字典名称
		DictType:    in.DictType,                                               // 字典类型
		State:       in.State,                                                  // 状态
		Remark:      sql.NullString{String: in.Remark, Valid: in.Remark != ""}, // 描述
		Sort:        in.Sort,                                                   // 排序
	})
	if err != nil {
		return nil, err
	}

	return &authenticationclient.CommonResp{}, nil
}

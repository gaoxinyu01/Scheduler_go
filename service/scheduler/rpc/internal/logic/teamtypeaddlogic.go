package logic

import (
	"Scheduler_go/service/scheduler/model"
	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type TeamTypeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTeamTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamTypeAddLogic {
	return &TeamTypeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 部门
func (l *TeamTypeAddLogic) TeamTypeAdd(in *schedulerclient.TeamTypeAddReq) (*schedulerclient.CommonResp, error) {
	_, err := l.svcCtx.TeamTypeModel.Insert(l.ctx, &model.TeamType{
		Id:          uuid.NewV4().String(),
		CreatedAt:   time.Now(),
		CreatedName: in.CreatedName,
		Name:        in.Name,
		Description: in.Description,
		TenantId:    in.TenantId,
	})

	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

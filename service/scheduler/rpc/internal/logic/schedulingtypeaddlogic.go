package logic

import (
	"Scheduler_go/service/scheduler/model"
	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingTypeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSchedulingTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingTypeAddLogic {
	return &SchedulingTypeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 排班类型
func (l *SchedulingTypeAddLogic) SchedulingTypeAdd(in *schedulerclient.SchedulingTypeAddReq) (*schedulerclient.CommonResp, error) {
	_, err := l.svcCtx.SchedulingTypeModel.Insert(l.ctx, &model.SchedulingType{
		Id:          uuid.NewV4().String(),
		CreatedAt:   time.Now(),
		CreatedName: in.CreatedName,
		Name:        in.Name,
		StartTime:   in.StartTime,
		EndTime:     in.EndTime,
		Remark:      sql.NullString{String: in.Remark, Valid: in.Remark != ""},
		Colour: sql.NullString{
			String: in.Colour,
			Valid:  in.Colour != "",
		},
		TenantId: in.TenantId,
	})
	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

package logic

import (
	"Scheduler_go/service/scheduler/model"
	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeamAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTeamAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamAddLogic {
	return &TeamAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 部门人员表
func (l *TeamAddLogic) TeamAdd(in *schedulerclient.TeamAddReq) (*schedulerclient.CommonResp, error) {
	// 查询是否存在这个部门
	_, err := l.svcCtx.TeamTypeModel.FindOne(l.ctx, in.TeamTypeId)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("没有该部门信息")
		}
		return nil, err
	}
	// 开启事务
	err = l.svcCtx.TeamModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {
		for _, v := range in.UserIds {
			// 查询用户是否已经在部门了
			res, err := l.svcCtx.TeamModel.FindOneByTeamTypeIdUserId("", v, in.TenantId)
			if err != nil {
				if err != sqlc.ErrNotFound {
					return err
				}
			}
			if err == nil {
				if res.Id != "" {
					return errors.New(fmt.Sprintf("%s,已存在该部门了", v))
				}
			}

			// 添加人员
			_, err = l.svcCtx.TeamModel.TransInsert(ctx, sqlx, &model.Team{
				Id:          uuid.NewV4().String(),
				CreatedAt:   time.Now(),
				CreatedName: in.CreatedName,
				UserId:      v,
				TeamTypeId:  in.TeamTypeId,
				TenantId:    in.TenantId,
			})
			if err != nil {
				return err
			}

		}
		return nil

	})
	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

package logic

import (
	"Scheduler_go/service/scheduler/model"
	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"
	"database/sql"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSchedulingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingAddLogic {
	return &SchedulingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 排班列表
func (l *SchedulingAddLogic) SchedulingAdd(in *schedulerclient.SchedulingAddReq) (*schedulerclient.CommonResp, error) {
	// 反序列化取数据
	var SchedulingAddList []model.SchedulingAddList
	err := jsonx.Unmarshal(in.Data, &SchedulingAddList)
	if err != nil {
		return nil, err
	}

	// 开启事务
	err = l.svcCtx.SchedulingModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {
		for _, item := range SchedulingAddList {
			// 查询部门信息
			res, err := l.svcCtx.TeamTypeModel.FindOne(l.ctx, item.TeamId)
			if err != nil {
				if err == sqlc.ErrNotFound {
					return errors.New(fmt.Sprintf("%s,没有该部门", item.TeamId))
				}
				return err
			}
			if res.TenantId != item.TenantId {
				return errors.New(fmt.Sprintf("%s,非法操作该租户没有该部门", item.TeamId))
			}

			// 添加用户到排班列表中
			for _, userData := range item.SchedulingUsers {
				// 查询当天内是否已存在排班,有排班无法排班
				netDayTime := item.Time + time.Hour.Milliseconds()*24
				count := l.svcCtx.SchedulingModel.FindOneByUserIdAndDayTime(userData.UserId, item.Time, netDayTime)
				if count > 0 {
					return errors.New(fmt.Sprintf("%s %s 已有排班 无法添加该天排班", userData.NickName, time.UnixMilli(item.Time)))
				}

				_, err = l.svcCtx.SchedulingModel.TransInsert(ctx, sqlx, &model.Scheduling{
					Id:          uuid.NewV4().String(),
					CreatedAt:   time.Now(),
					CreatedName: item.CreatedName,
					Time:        item.Time,
					Name:        item.Name,
					StartTime:   item.StartTime,
					EndTime:     item.EndTime,
					TeamName:    res.Name,
					UserName:    userData.NickName,
					Colour: sql.NullString{
						String: item.Colour,
						Valid:  item.Colour != "",
					},
					TeamId:     item.TeamId,
					UserId:     userData.UserId,
					TenantId:   item.TenantId,
					IsFinished: 0,
				})
				if err != nil {
					return err
				}

			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

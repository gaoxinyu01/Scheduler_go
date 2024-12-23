package scheduling

import (
	"Scheduler_go/common"
	"Scheduler_go/common/global/jwtx"
	"Scheduler_go/common/msg"
	"Scheduler_go/service/manage/authentication/authenticationclient"
	"Scheduler_go/service/scheduler/model"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchedulingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSchedulingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulingAddLogic {
	return &SchedulingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SchedulingAddLogic) SchedulingAdd(req *types.SchedulingAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	var SchedulingAddLists []model.SchedulingAddList

	// 组装数据
	for _, item := range req.Data {
		var schedulingUsers []model.SchedulingUser
		for _, v := range item.UserIds {
			res, err := l.svcCtx.AuthenticationRpc.SysUserFindOne(l.ctx, &authenticationclient.SysUserFindOneReq{
				Id: v,
			})
			if err != nil {
				return nil, common.NewDefaultError(err.Error())
			}
			schedulingUsers = append(schedulingUsers, model.SchedulingUser{
				UserId:   res.Id,
				NickName: res.NickName,
			})
		}
		SchedulingAddLists = append(SchedulingAddLists, model.SchedulingAddList{
			CreatedName:     tokenData.NickName,
			Time:            item.Time,
			Name:            item.Name,
			StartTime:       item.StartTime,
			EndTime:         item.EndTime,
			TeamId:          item.TeamId,
			SchedulingUsers: schedulingUsers,
			Colour:          item.Colour,
			TenantId:        tokenData.TenantId,
		})
	}

	data, err := jsonx.Marshal(SchedulingAddLists)
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	_, err = l.svcCtx.SchedulerRpc.SchedulingAdd(l.ctx, &schedulerclient.SchedulingAddReq{
		Data: data,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}

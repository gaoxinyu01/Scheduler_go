package logic

import (
	"context"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeamFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTeamFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeamFindListLogic {
	return &TeamFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TeamFindListLogic) TeamFindList(in *schedulerclient.TeamFindListReq) (*schedulerclient.TeamFindListResp, error) {
	/*
	  int64  current = 1;
	  int64 page_size = 2;
	  string nick_name = 3; // 昵称
	  string major = 4; // 专业
	  string position = 5; // 岗位
	  string telephone =6; // 手机号
	  string team_type_id = 7; // 部门ID
	  string tenant_id = 8; // 租户ID
	*/
	all, err := l.svcCtx.TeamModel.FindList(in.Current, in.PageSize, in.NickName,
		in.Major, in.Position, in.Telephone, in.TeamTypeId, in.TenantId)

	if err != nil {
		return nil, err
	}

	var TeamUser []*schedulerclient.TeamUser

	for _, item := range *all {
		TeamUser = append(TeamUser, &schedulerclient.TeamUser{
			Id:        item.Id,
			Account:   item.Account,
			NickName:  item.NickName,
			Major:     item.Major.String,
			Position:  item.Position.String,
			Avatar:    item.Avatar.String,
			Email:     item.Email,
			Telephone: item.Telephone,
			State:     item.State,
			UserId:    item.Uid,
		})

	}

	total := l.svcCtx.TeamModel.Count(in.NickName,
		in.Major, in.Position, in.Telephone, in.TeamTypeId, in.TenantId)
	return &schedulerclient.TeamFindListResp{
		Total: total,
		List:  TeamUser,
	}, nil
}

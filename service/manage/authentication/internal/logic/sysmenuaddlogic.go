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

type SysMenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuAddLogic {
	return &SysMenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 菜单
func (l *SysMenuAddLogic) SysMenuAdd(in *authenticationclient.SysMenuAddReq) (*authenticationclient.CommonResp, error) {

	_, err := l.svcCtx.SysMenuModel.Insert(l.ctx, &model.SysMenu{
		CreatedAt:   time.Now(),                                                    // 创建时间
		MenuType:    in.MenuType,                                                   // 菜单类型(层级关系)
		Name:        in.Name,                                                       // 菜单名称
		Title:       in.Title,                                                      // 标题
		Path:        in.Path,                                                       // 路径
		Component:   in.Component,                                                  // 本地路径
		Redirect:    sql.NullString{String: in.Redirect, Valid: in.Redirect != ""}, // 跳转
		Sort:        in.Sort,                                                       // sort
		Icon:        sql.NullString{String: in.Icon, Valid: in.Icon != ""},         // 图标
		IsHide:      in.IsHide,                                                     // 是否隐藏
		IsKeepAlive: in.IsKeepAlive,                                                // 是否缓存
		ParentId:    in.ParentId,                                                   // 父ID
		IsHome:      in.IsHome,                                                     // 是否首页
		IsMain:      in.IsMain,                                                     // 是否主菜单
		CreatedName: in.CreatedName,                                                // 创建人
	})
	if err != nil {
		return nil, err
	}

	return &authenticationclient.CommonResp{}, nil
}

// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: archive.proto

package server

import (
	"context"

	"Scheduler_go/service/manage/archive/rpc/archiveclient"
	"Scheduler_go/service/manage/archive/rpc/internal/logic"
	"Scheduler_go/service/manage/archive/rpc/internal/svc"
)

type ArchiveServer struct {
	svcCtx *svc.ServiceContext
	archiveclient.UnimplementedArchiveServer
}

func NewArchiveServer(svcCtx *svc.ServiceContext) *ArchiveServer {
	return &ArchiveServer{
		svcCtx: svcCtx,
	}
}

// 用户日志
func (s *ArchiveServer) AppLoggerAdd(ctx context.Context, in *archiveclient.AppLoggerAddReq) (*archiveclient.CommonResp, error) {
	l := logic.NewAppLoggerAddLogic(ctx, s.svcCtx)
	return l.AppLoggerAdd(in)
}

func (s *ArchiveServer) AppLoggerFindList(ctx context.Context, in *archiveclient.AppLoggerFindListReq) (*archiveclient.AppLoggerFindListResp, error) {
	l := logic.NewAppLoggerFindListLogic(ctx, s.svcCtx)
	return l.AppLoggerFindList(in)
}

package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"Scheduler_go/service/scheduler/rpc/internal/svc"
	"Scheduler_go/service/scheduler/rpc/schedulerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttendanceDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttendanceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttendanceDeleteLogic {
	return &AttendanceDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AttendanceDeleteLogic) AttendanceDelete(in *schedulerclient.AttendanceDeleteReq) (*schedulerclient.CommonResp, error) {
	res, err := l.svcCtx.AttendanceModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("Attendance没有该ID：" + in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, errors.New("Attendance该ID已被删除：" + in.Id)
	}
	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作")
	}

	res.DeletedAt.Time = time.Now()
	res.DeletedAt.Valid = true
	res.DeletedName.String = in.DeletedName
	res.DeletedName.Valid = true

	err = l.svcCtx.AttendanceModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}

	return &schedulerclient.CommonResp{}, nil
}

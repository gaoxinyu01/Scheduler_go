package svc

import (
	"Scheduler_go/service/scheduler/model"
	"Scheduler_go/service/scheduler/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	//打卡
	AttendanceModel model.AttendanceModel

	// 部门
	TeamTypeModel model.TeamTypeModel

	// 部门人员
	TeamModel model.TeamModel

	// 排班类型
	SchedulingTypeModel model.SchedulingTypeModel

	// 排班列表
	SchedulingModel model.SchedulingModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		AttendanceModel:     model.NewAttendanceModel(conn, c.CacheRedis),
		TeamTypeModel:       model.NewTeamTypeModel(conn, c.CacheRedis),
		TeamModel:           model.NewTeamModel(conn, c.CacheRedis),
		SchedulingTypeModel: model.NewSchedulingTypeModel(conn, c.CacheRedis),
		SchedulingModel:     model.NewSchedulingModel(conn, c.CacheRedis),
	}
}

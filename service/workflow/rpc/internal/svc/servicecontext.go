package svc

import (
	"Scheduler_go/service/workflow/model"
	"Scheduler_go/service/workflow/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	//流程定义
	ProcDefModel model.ProcDefModel
	//流程定义
	HistProcDefModel model.HistProcDefModel
	//流程实例
	ProcInstModel model.ProcInstModel
	//执行
	ProcExecutionModel model.ProcExecutionModel
	//历史执行
	HistProcExecutionModel model.HistProcExecutionModel
	//流程实例变量
	ProcInstVariableModel model.ProcInstVariableModel
	//任务
	ProcTaskModel model.ProcTaskModel
	//历史任务
	HistProcTaskModel model.HistProcTaskModel

	HistProcInstModel model.HistProcInstModel

	HistProcInstVariableModel model.HistProcInstVariableModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                    c,
		ProcDefModel:              model.NewProcDefModel(conn, c.CacheRedis),
		HistProcDefModel:          model.NewHistProcDefModel(conn, c.CacheRedis),
		ProcExecutionModel:        model.NewProcExecutionModel(conn, c.CacheRedis),
		HistProcExecutionModel:    model.NewHistProcExecutionModel(conn, c.CacheRedis),
		ProcInstModel:             model.NewProcInstModel(conn, c.CacheRedis),
		ProcInstVariableModel:     model.NewProcInstVariableModel(conn, c.CacheRedis),
		ProcTaskModel:             model.NewProcTaskModel(conn, c.CacheRedis),
		HistProcTaskModel:         model.NewHistProcTaskModel(conn, c.CacheRedis),
		HistProcInstModel:         model.NewHistProcInstModel(conn, c.CacheRedis),
		HistProcInstVariableModel: model.NewHistProcInstVariableModel(conn, c.CacheRedis),
	}
}

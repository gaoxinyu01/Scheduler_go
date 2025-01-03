// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package handler

import (
	"net/http"

	attendance "Scheduler_go/service/scheduler/api/internal/handler/attendance"
	scheduling "Scheduler_go/service/scheduler/api/internal/handler/scheduling"
	schedulingType "Scheduler_go/service/scheduler/api/internal/handler/schedulingType"
	team "Scheduler_go/service/scheduler/api/internal/handler/team"
	teamType "Scheduler_go/service/scheduler/api/internal/handler/teamType"
	"Scheduler_go/service/scheduler/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/scheduler/Attendance",
				Handler: attendance.AttendanceAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/scheduler/Attendance",
				Handler: attendance.AttendanceUpHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scheduler/Attendance",
				Handler: attendance.AttendanceListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/scheduler/Attendance/:id",
				Handler: attendance.AttendanceDelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scheduler/AttendanceInfo",
				Handler: attendance.AttendanceInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/scheduler/AttendancePatch",
				Handler: attendance.AttendancePatchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scheduler/DayAttendanceByTimeInfo",
				Handler: attendance.DayAttendanceByTimeInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scheduler/DayAttendanceInfo",
				Handler: attendance.DayAttendanceInfoHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/scheduler/scheduling",
				Handler: scheduling.SchedulingAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/scheduler/scheduling",
				Handler: scheduling.SchedulingUpHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scheduler/scheduling",
				Handler: scheduling.SchedulingListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/scheduler/scheduling/:id",
				Handler: scheduling.SchedulingDelHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/scheduler/schedulingType",
				Handler: schedulingType.SchedulingTypeAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/scheduler/schedulingType",
				Handler: schedulingType.SchedulingTypeUpHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scheduler/schedulingType",
				Handler: schedulingType.SchedulingTypeListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/scheduler/schedulingType/:id",
				Handler: schedulingType.SchedulingTypeDelHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/scheduler/notTeamUser",
				Handler: team.NotTeamUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/scheduler/team",
				Handler: team.TeamAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scheduler/team",
				Handler: team.TeamListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/scheduler/team/:id",
				Handler: team.TeamDelHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/scheduler/teamType",
				Handler: teamType.TeamTypeAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/scheduler/teamType",
				Handler: teamType.TeamTypeUpHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scheduler/teamType",
				Handler: teamType.TeamTypeListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/scheduler/teamType/:id",
				Handler: teamType.TeamTypeDelHandler(serverCtx),
			},
		},
	)
}

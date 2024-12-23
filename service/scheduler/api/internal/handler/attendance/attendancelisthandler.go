package attendance

import (
	"Scheduler_go/common/responsex"
	"Scheduler_go/service/scheduler/api/internal/logic/attendance"
	"Scheduler_go/service/scheduler/api/internal/svc"
	"Scheduler_go/service/scheduler/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"Scheduler_go/common"
)

func AttendanceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AttendanceListRequest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := attendance.NewAttendanceListLogic(r.Context(), svcCtx)
		resp, err := l.AttendanceList(&req)
		responsex.HttpResult(r, w, req, resp, err, svcCtx.ArchiveRpc)
	}
}
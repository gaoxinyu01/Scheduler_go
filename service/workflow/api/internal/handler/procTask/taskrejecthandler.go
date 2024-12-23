package procTask

import (
	"Scheduler_go/common/responsex"
	"Scheduler_go/service/workflow/api/internal/logic/procTask"
	"Scheduler_go/service/workflow/api/internal/svc"
	"Scheduler_go/service/workflow/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"Scheduler_go/common"
)

func TaskRejectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskRejectRequest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := procTask.NewTaskRejectLogic(r.Context(), svcCtx)
		resp, err := l.TaskReject(&req)
		responsex.HttpResult(r, w, req, resp, err, svcCtx.ArchiveRpc)
	}
}

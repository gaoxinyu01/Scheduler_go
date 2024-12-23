package procDef

import (
	"Scheduler_go/common/responsex"
	"Scheduler_go/service/workflow/api/internal/logic/procDef"
	"Scheduler_go/service/workflow/api/internal/svc"
	"Scheduler_go/service/workflow/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"Scheduler_go/common"
)

func ProcDefListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProcDefListRequest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := procDef.NewProcDefListLogic(r.Context(), svcCtx)
		resp, err := l.ProcDefList(&req)
		responsex.HttpResult(r, w, req, resp, err, svcCtx.ArchiveRpc)
	}
}
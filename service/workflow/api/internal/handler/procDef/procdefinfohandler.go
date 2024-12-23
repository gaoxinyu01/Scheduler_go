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

func ProcDefInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProcDefInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := procDef.NewProcDefInfoLogic(r.Context(), svcCtx)
		resp, err := l.ProcDefInfo(&req)
		responsex.HttpResult(r, w, req, resp, err, svcCtx.ArchiveRpc)
	}
}

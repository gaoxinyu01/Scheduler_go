package appLog

import (
	"Scheduler_go/common"
	"Scheduler_go/common/responsex"
	"Scheduler_go/service/manage/archive/api/internal/logic/appLog"
	"Scheduler_go/service/manage/archive/api/internal/svc"
	"Scheduler_go/service/manage/archive/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AppLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppLogReqest
		if err := httpx.Parse(r, &req); err != nil {
			responsex.HttpResult(r, w, req, "", common.NewParamError(err.Error()), svcCtx.ArchiveRpc)
			return
		}

		l := appLog.NewAppLogLogic(r.Context(), svcCtx)
		resp, err := l.AppLog(&req)
		responsex.HttpResult(r, w, req, resp, err, nil)
	}
}

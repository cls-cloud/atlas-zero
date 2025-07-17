package operLog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"monitor/internal/logic/operLog"
	"monitor/internal/svc"
	"monitor/internal/types"
)

func PageSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageSetOperLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := operLog.NewPageSetLogic(r.Context(), svcCtx)
		resp, err := l.PageSet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

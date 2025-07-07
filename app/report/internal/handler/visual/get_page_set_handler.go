package visual

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"report/internal/logic/visual"
	"report/internal/svc"
	"report/internal/types"
)

func GetPageSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VisualPageSetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := visual.NewGetPageSetLogic(r.Context(), svcCtx)
		resp, err := l.GetPageSet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

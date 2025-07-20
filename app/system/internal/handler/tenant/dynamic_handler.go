package tenant

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"system/internal/logic/tenant"
	"system/internal/svc"
	"system/internal/types"
)

func DynamicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tenant.NewDynamicLogic(r.Context(), svcCtx)
		err := l.Dynamic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}

package health

import (
	"net/http"
	"ovra/app/demo/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJsonCtx(r.Context(), w, map[string]string{"status": "ok"})
	}
}

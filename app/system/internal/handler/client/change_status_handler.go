package client

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"system/internal/logic/client"
	"system/internal/svc"
	"system/internal/types"
)

func ChangeStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClientQuery
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := client.NewChangeStatusLogic(r.Context(), svcCtx)
		resp, err := l.ChangeStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

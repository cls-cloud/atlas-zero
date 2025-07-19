package oss

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"resource/internal/logic/oss"
	"resource/internal/svc"
	"resource/internal/types"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oss.NewDeleteLogic(r.Context(), svcCtx)
		err := l.Delete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}

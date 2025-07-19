package oss

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"resource/internal/logic/oss"
	"resource/internal/svc"
	"resource/internal/types"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyOssReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oss.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}

package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"system/internal/logic/user"
	"system/internal/svc"
	"system/internal/types"
)

func ResetPwdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetPwdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewResetPwdLogic(r.Context(), svcCtx)
		err := l.ResetPwd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}

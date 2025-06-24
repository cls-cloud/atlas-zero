package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("请求前...")
		next(w, r)
		logx.Info("请求后....")
	}
}

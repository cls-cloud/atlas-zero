package middleware

import (
	"net/http"
	"system/internal/config"
)

func ApiMiddleware(cfg config.Config) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if cfg.RestConf.Mode == "pre" && isUpdateMethod(r.Method) {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte(`{"code":500, "msg":"演示模式, 不允许操作"}`))
				if err != nil {
					return
				}
				return
			}
			next(w, r)
		}
	}
}

func isUpdateMethod(method string) bool {
	switch method {
	case http.MethodPut, http.MethodDelete:
		return true
	default:
		return false
	}
}

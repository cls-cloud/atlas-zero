package middleware

import (
	"net/http"
)

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//指定允许其他域名访问
		//ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Origin", "*") //跨域：CORS(跨来源资源共享)策略
		//预检结果缓存时间
		w.Header().Set("Access-Control-Max-Age", "86400")
		//允许的请求类型（GET,POST等）
		w.Header().Set("Access-Control-Allow-Methods", "*")
		//允许的请求头字段
		w.Header().Set("Access-Control-Allow-Headers", "*")
		//是否允许后续请求携带认证信息（cookies）,该值只能是true,否则不返回
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

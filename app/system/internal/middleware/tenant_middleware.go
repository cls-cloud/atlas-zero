package middleware

import (
	"context"
	"net/http"
)

// TenantContextKey 用于存放 tenant_id
type TenantContextKey struct{}

// TenantMiddleware 中间件：从请求中提取 tenant_id
func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取租户ID，例如从请求头中读取
		tenantID := r.Header.Get("X-Tenant-ID")
		if tenantID == "" {
			http.Error(w, "Missing tenant_id", http.StatusBadRequest)
			return
		}

		// 将 tenant_id 存入上下文
		ctx := context.WithValue(r.Context(), TenantContextKey{}, tenantID)
		// 使用新的上下文继续执行下一个处理器
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetTenantID 从上下文获取 tenant_id
func GetTenantID(ctx context.Context) string {
	if tenantID, ok := ctx.Value(TenantContextKey{}).(string); ok {
		return tenantID
	}
	return ""
}

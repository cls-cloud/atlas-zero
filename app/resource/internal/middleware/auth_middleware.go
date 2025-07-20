package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"resource/internal/config"
	"strings"
	"toolkit/auth"
	"toolkit/tenant"
)

type AuthMiddleware struct {
	c   config.Config
	rds *redis.Redis
}

func NewAuthMiddleware(c config.Config, rds *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{c: c, rds: rds}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return m.ExecHandle(next)
}

func (m *AuthMiddleware) ExecHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authorization, "Bearer ")
		uc, err := auth.AnalyseToken(tokenString, m.c.JwtAuth.AccessSecret)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		authInstance := auth.NewAuth(m.rds, &uc.UserInfo)
		key := fmt.Sprintf(auth.TokenKey, uc.ClientId, uc.UserId)

		expired, err := authInstance.CheckToken(r.Context(), key)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		if expired {
			http.Error(w, "Unauthorized: token expired (idle timeout)", http.StatusUnauthorized)
			return
		}

		tenantId, err := tenant.GetTenantId(r.Context(), m.rds, &uc.UserInfo)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		r.Header.Set(auth.UserIDKey, uc.UserId)
		r.Header.Set(auth.TenantIDKey, tenantId)
		r.Header.Set(auth.ClientIDKey, uc.ClientId)
		ctx := context.WithValue(r.Context(), auth.UserIDKey, uc.UserId)
		ctx = context.WithValue(ctx, auth.TenantIDKey, tenantId)
		ctx = context.WithValue(ctx, auth.ClientIDKey, uc.ClientId)
		next(w, r.WithContext(ctx))
	}
}

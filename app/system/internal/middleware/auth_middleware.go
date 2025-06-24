package middleware

import (
	"context"
	"net/http"
	"strings"
	"system/internal/config"
	"toolkit/helper"
)

type AuthMiddleware struct {
	c config.Config
}

func NewAuthMiddleware(c config.Config) *AuthMiddleware {
	return &AuthMiddleware{c: c}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		uc, err := helper.AnalyseToken(tokenString, m.c.JwtAuth.AccessSecret)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		r.Header.Set("UserId", uc.UserId)
		r.Header.Set("TenantId", uc.TenantId)
		ctx := context.WithValue(r.Context(), "user_id", uc.UserId)
		ctx = context.WithValue(ctx, "tenant_id", uc.TenantId)
		next(w, r.WithContext(ctx))
	}
}

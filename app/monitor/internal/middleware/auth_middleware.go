package middleware

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"monitor/internal/config"
	"net/http"
	"toolkit/middlewares"
)

type AuthMiddleware struct {
	c   config.Config
	rds *redis.Redis
}

func NewAuthMiddleware(c config.Config, rds *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{c: c, rds: rds}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middlewares.ExecHandle(next, m.c.JwtAuth.AccessSecret, m.rds, m.c.JwtAuth.MultipleLoginDevices)
}

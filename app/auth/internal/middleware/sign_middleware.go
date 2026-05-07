package middleware

import (
	"net/http"
	"ovra/app/auth/internal/config"
	"ovra/toolkit/middlewares"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SignMiddleware struct {
	c   config.Config
	rds *redis.Redis
}

func NewSignMiddleware(c config.Config, rds *redis.Redis) *SignMiddleware {
	return &SignMiddleware{c: c, rds: rds}
}

func (m *SignMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	if !m.c.Sign.Enabled {
		return next
	}

	return middlewares.SignExecHandle(next, m.rds)
}

package middleware

import (
	"net/http"
	"ovra/toolkit/middlewares"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SignMiddleware struct {
	rds *redis.Redis
}

func NewSignMiddleware(rds *redis.Redis) *SignMiddleware {
	return &SignMiddleware{rds: rds}
}

func (m *SignMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middlewares.SignExecHandle(next, m.rds)
}

package middleware

import (
	"net/http"
	"ovra/app/system/internal/config"
	"ovra/toolkit/middlewares"
)

type ApiDecryptMiddleware struct {
	c config.Config
}

func NewApiDecryptMiddleware(c config.Config) *ApiDecryptMiddleware {
	return &ApiDecryptMiddleware{c: c}
}

func (m *ApiDecryptMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middlewares.ApiEncryptMiddleware(m.c.ApiDecrypt)(next)
}

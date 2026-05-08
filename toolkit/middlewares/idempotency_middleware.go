package middlewares

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"

	"ovra/toolkit/auth"
	"ovra/toolkit/errx"
	"ovra/toolkit/helper"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	defaultIdempotencyHeader  = "Idempotency-Key"
	defaultIdempotencyExpire  = 10
	defaultIdempotencyMaxBody = 1 << 20
)

type IdempotencyConfig struct {
	Enabled       bool
	Header        string
	ExpireSeconds int
	IncludePaths  []string
	ExcludePaths  []string
}

func IdempotencyMiddleware(rds *redis.Redis, c IdempotencyConfig) func(http.HandlerFunc) http.HandlerFunc {
	header := c.Header
	if header == "" {
		header = defaultIdempotencyHeader
	}
	expire := c.ExpireSeconds
	if expire <= 0 {
		expire = defaultIdempotencyExpire
	}

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !c.Enabled || rds == nil || !shouldCheckIdempotency(r, c) {
				next(w, r)
				return
			}

			key, ok := buildIdempotencyKey(r, header)
			if !ok {
				next(w, r)
				return
			}

			locked, err := rds.SetnxExCtx(r.Context(), key, "1", expire)
			if err != nil {
				logx.WithContext(r.Context()).Errorf("idempotency setnx failed: %v", err)
				httpx.OkJsonCtx(r.Context(), w, helper.Fail(errx.New(errx.CodeInternal, "请求幂等校验失败")))
				return
			}
			if !locked {
				httpx.OkJsonCtx(r.Context(), w, helper.Fail(errx.New(errx.CodeBizErr, "重复请求，请稍后再试")))
				return
			}

			next(w, r)
		}
	}
}

func shouldCheckIdempotency(r *http.Request, c IdempotencyConfig) bool {
	if matchPath(r.URL.Path, c.ExcludePaths) {
		return false
	}
	if matchPath(r.URL.Path, c.IncludePaths) {
		return true
	}

	switch r.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
		return true
	default:
		return false
	}
}

func matchPath(path string, patterns []string) bool {
	for _, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}
		if strings.HasSuffix(pattern, "*") {
			if strings.HasPrefix(path, strings.TrimSuffix(pattern, "*")) {
				return true
			}
			continue
		}
		if path == pattern {
			return true
		}
	}
	return false
}

func buildIdempotencyKey(r *http.Request, header string) (string, bool) {
	if idempotencyKey := strings.TrimSpace(r.Header.Get(header)); idempotencyKey != "" {
		return hashIdempotencyKey("header", requestOwner(r), r.Method, r.URL.Path, idempotencyKey), true
	}

	if strings.HasPrefix(strings.ToLower(r.Header.Get("Content-Type")), "multipart/") {
		return "", false
	}

	body, ok := readRequestBody(r, defaultIdempotencyMaxBody)
	if !ok {
		return "", false
	}
	return hashIdempotencyKey("fingerprint", requestOwner(r), r.Method, r.URL.RequestURI(), string(body)), true
}

func readRequestBody(r *http.Request, maxBody int64) ([]byte, bool) {
	if r.Body == nil {
		return nil, true
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, maxBody+1))
	if err != nil {
		return nil, false
	}
	r.Body = io.NopCloser(bytes.NewReader(body))
	if int64(len(body)) > maxBody {
		return nil, false
	}
	return body, true
}

func requestOwner(r *http.Request) string {
	tenantId := valueFromRequest(r, auth.TenantIDKey)
	userId := valueFromRequest(r, auth.UserIDKey)
	clientId := valueFromRequest(r, auth.ClientIDKey)
	if tenantId != "" || userId != "" || clientId != "" {
		return tenantId + ":" + userId + ":" + clientId
	}
	if authorization := strings.TrimSpace(r.Header.Get("Authorization")); authorization != "" {
		return authorization
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}
	return r.RemoteAddr
}

func valueFromRequest(r *http.Request, key string) string {
	if value := r.Header.Get(key); value != "" {
		return value
	}
	if value := r.Context().Value(key); value != nil {
		switch v := value.(type) {
		case string:
			return v
		case int64:
			return strconv.FormatInt(v, 10)
		case int:
			return strconv.Itoa(v)
		}
	}
	return ""
}

func hashIdempotencyKey(parts ...string) string {
	h := sha256.New()
	for _, part := range parts {
		_, _ = h.Write([]byte(part))
		_, _ = h.Write([]byte{0})
	}
	return "idempotency:" + hex.EncodeToString(h.Sum(nil))
}

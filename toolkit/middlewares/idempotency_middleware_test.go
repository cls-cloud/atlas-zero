package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/stores/redis/redistest"
)

func TestIdempotencyMiddlewareBlocksDuplicateHeader(t *testing.T) {
	rds, clean := redistest.CreateRedisWithClean(t)
	defer clean()

	middleware := IdempotencyMiddleware(rds, IdempotencyConfig{
		Enabled:       true,
		ExpireSeconds: 60,
	})
	var calls int
	handler := middleware(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
	})

	first := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"sku":1}`))
	first.Header.Set(defaultIdempotencyHeader, "same-key")
	first.Header.Set("X-User-Id", "100")
	handler(httptest.NewRecorder(), first)

	second := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"sku":2}`))
	second.Header.Set(defaultIdempotencyHeader, "same-key")
	second.Header.Set("X-User-Id", "100")
	recorder := httptest.NewRecorder()
	handler(recorder, second)

	require.Equal(t, 1, calls)
	require.Contains(t, recorder.Body.String(), "重复请求")
}

func TestIdempotencyMiddlewareRestoresRequestBody(t *testing.T) {
	rds, clean := redistest.CreateRedisWithClean(t)
	defer clean()

	middleware := IdempotencyMiddleware(rds, IdempotencyConfig{
		Enabled:       true,
		ExpireSeconds: 60,
	})
	handler := middleware(func(w http.ResponseWriter, r *http.Request) {
		body := make([]byte, r.ContentLength)
		_, err := r.Body.Read(body)
		require.NoError(t, err)
		_, _ = w.Write(body)
	})

	req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"sku":1}`))
	recorder := httptest.NewRecorder()
	handler(recorder, req)

	require.Equal(t, `{"sku":1}`, recorder.Body.String())
}

func TestIdempotencyMiddlewareSkipsGet(t *testing.T) {
	rds, clean := redistest.CreateRedisWithClean(t)
	defer clean()

	middleware := IdempotencyMiddleware(rds, IdempotencyConfig{
		Enabled:       true,
		ExpireSeconds: 60,
	})
	var calls int
	handler := middleware(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	req.Header.Set(defaultIdempotencyHeader, "same-key")
	handler(httptest.NewRecorder(), req)
	handler(httptest.NewRecorder(), req)

	require.Equal(t, 2, calls)
}

func TestIdempotencyMiddlewareIncludesConfiguredGetPath(t *testing.T) {
	rds, clean := redistest.CreateRedisWithClean(t)
	defer clean()

	middleware := IdempotencyMiddleware(rds, IdempotencyConfig{
		Enabled:       true,
		ExpireSeconds: 60,
		IncludePaths:  []string{"/auth/code"},
	})
	var calls int
	handler := middleware(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/auth/code", nil)
	req.RemoteAddr = "127.0.0.1:12345"
	handler(httptest.NewRecorder(), req)

	recorder := httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/auth/code", nil)
	req.RemoteAddr = "127.0.0.1:12345"
	handler(recorder, req)

	require.Equal(t, 1, calls)
	require.Contains(t, recorder.Body.String(), "重复请求")
}

func TestIdempotencyMiddlewareExcludesConfiguredPath(t *testing.T) {
	rds, clean := redistest.CreateRedisWithClean(t)
	defer clean()

	middleware := IdempotencyMiddleware(rds, IdempotencyConfig{
		Enabled:       true,
		ExpireSeconds: 60,
		ExcludePaths:  []string{"/orders"},
	})
	var calls int
	handler := middleware(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"sku":1}`))
	handler(httptest.NewRecorder(), req)
	req = httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"sku":1}`))
	handler(httptest.NewRecorder(), req)

	require.Equal(t, 2, calls)
}

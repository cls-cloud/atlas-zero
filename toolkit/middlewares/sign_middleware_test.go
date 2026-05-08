package middlewares

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/stores/redis/redistest"
)

func TestSignExecHandleVerifiesQueryAndBody(t *testing.T) {
	rds, clean := redistest.CreateRedisWithClean(t)
	defer clean()
	require.NoError(t, rds.Set("sign:secret:test-key", "test-secret"))

	timestamp := time.Now().Unix()
	body := `{"amount":100}`
	sign := Sign(SignParams{
		Method:     http.MethodPost,
		Path:       "/pay",
		Query:      url.Values{"b": []string{"2"}, "a": []string{"1"}},
		AccessKey:  "test-key",
		Timestamp:  timestamp,
		Nonce:      "nonce-1",
		BodySHA256: SHA256Hex([]byte(body)),
	}.CanonicalString(), "test-secret")

	var calls int
	handler := SignExecHandle(func(w http.ResponseWriter, r *http.Request) {
		calls++
		readBody := make([]byte, r.ContentLength)
		_, err := r.Body.Read(readBody)
		require.NoError(t, err)
		require.Equal(t, body, string(readBody))
		w.WriteHeader(http.StatusOK)
	}, rds)

	req := httptest.NewRequest(http.MethodPost, "/pay?b=2&a=1", strings.NewReader(body))
	req.Header.Set(signAccessKeyHeader, "test-key")
	req.Header.Set(signTimestampHeader, strconv.FormatInt(timestamp, 10))
	req.Header.Set(signNonceHeader, "nonce-1")
	req.Header.Set(signHeader, sign)
	handler(httptest.NewRecorder(), req)

	require.Equal(t, 1, calls)
}

func TestSignExecHandleRejectsBodyTampering(t *testing.T) {
	rds, clean := redistest.CreateRedisWithClean(t)
	defer clean()
	require.NoError(t, rds.Set("sign:secret:test-key", "test-secret"))

	timestamp := time.Now().Unix()
	sign := Sign(SignParams{
		Method:     http.MethodPost,
		Path:       "/pay",
		Query:      url.Values{},
		AccessKey:  "test-key",
		Timestamp:  timestamp,
		Nonce:      "nonce-2",
		BodySHA256: SHA256Hex([]byte(`{"amount":100}`)),
	}.CanonicalString(), "test-secret")

	handler := SignExecHandle(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("handler should not be called")
	}, rds)

	req := httptest.NewRequest(http.MethodPost, "/pay", strings.NewReader(`{"amount":999}`))
	req.Header.Set(signAccessKeyHeader, "test-key")
	req.Header.Set(signTimestampHeader, strconv.FormatInt(timestamp, 10))
	req.Header.Set(signNonceHeader, "nonce-2")
	req.Header.Set(signHeader, sign)
	recorder := httptest.NewRecorder()
	handler(recorder, req)

	require.Contains(t, recorder.Body.String(), "signature mismatch")
}

func TestSignExecHandleRejectsNonceReplay(t *testing.T) {
	rds, clean := redistest.CreateRedisWithClean(t)
	defer clean()
	require.NoError(t, rds.Set("sign:secret:test-key", "test-secret"))

	timestamp := time.Now().Unix()
	body := ""
	sign := Sign(SignParams{
		Method:     http.MethodGet,
		Path:       "/pay",
		Query:      url.Values{},
		AccessKey:  "test-key",
		Timestamp:  timestamp,
		Nonce:      "nonce-3",
		BodySHA256: SHA256Hex([]byte(body)),
	}.CanonicalString(), "test-secret")

	var calls int
	handler := SignExecHandle(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
	}, rds)

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(http.MethodGet, "/pay", nil)
		req.Header.Set(signAccessKeyHeader, "test-key")
		req.Header.Set(signTimestampHeader, strconv.FormatInt(timestamp, 10))
		req.Header.Set(signNonceHeader, "nonce-3")
		req.Header.Set(signHeader, sign)
		recorder := httptest.NewRecorder()
		handler(recorder, req)
		if i == 1 {
			require.Contains(t, recorder.Body.String(), "signature replayed")
		}
	}

	require.Equal(t, 1, calls)
}

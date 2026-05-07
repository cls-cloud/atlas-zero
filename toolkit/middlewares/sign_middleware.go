package middlewares

import (
	"bytes"
	"crypto/hmac"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"ovra/toolkit/errx"
	"ovra/toolkit/helper"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	signAccessKeyHeader = "X-Key"
	signTimestampHeader = "X-Timestamp"
	signNonceHeader     = "X-Nonce"
	signHeader          = "X-Sign"
	signWindowSeconds   = 300
	signMaxBodyBytes    = 10 << 20
)

func SignExecHandle(next http.HandlerFunc, rds *redis.Redis) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessKey := strings.TrimSpace(r.Header.Get(signAccessKeyHeader))
		timestampStr := strings.TrimSpace(r.Header.Get(signTimestampHeader))
		nonce := strings.TrimSpace(r.Header.Get(signNonceHeader))
		sign := strings.TrimSpace(r.Header.Get(signHeader))
		if accessKey == "" || timestampStr == "" || nonce == "" || sign == "" {
			writeSignError(w, r, "missing sign headers")
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			writeSignError(w, r, "invalid timestamp")
			return
		}
		now := time.Now().Unix()
		if abs(now-timestamp) > signWindowSeconds {
			writeSignError(w, r, "timestamp expired")
			return
		}

		secret, err := getSecretByAccessKey(rds, accessKey)
		if err != nil {
			writeSignError(w, r, err.Error())
			return
		}

		body, ok := readSignBody(r, signMaxBodyBytes)
		if !ok {
			writeSignError(w, r, "invalid request body")
			return
		}

		canonical := SignParams{
			Method:     r.Method,
			Path:       r.URL.Path,
			Query:      r.URL.Query(),
			AccessKey:  accessKey,
			Timestamp:  timestamp,
			Nonce:      nonce,
			BodySHA256: SHA256Hex(body),
		}.CanonicalString()

		expectSign := Sign(canonical, secret)
		if !hmac.Equal([]byte(expectSign), []byte(sign)) {
			writeSignError(w, r, "signature mismatch")
			return
		}

		locked, err := markNonceUsed(rds, r, accessKey, nonce)
		if err != nil {
			writeSignError(w, r, "signature nonce check failed")
			return
		}
		if !locked {
			writeSignError(w, r, "signature replayed")
			return
		}

		next(w, r)
	}
}

func getSecretByAccessKey(rds *redis.Redis, accessKey string) (string, error) {
	if rds == nil {
		return "", errors.New("sign redis not configured")
	}

	key := "sign:secret:" + accessKey
	val, err := rds.Get(key)
	if err != nil {
		return "", err
	}
	if val == "" {
		return "", errors.New("access key not found")
	}
	return val, nil
}

func markNonceUsed(rds *redis.Redis, r *http.Request, accessKey, nonce string) (bool, error) {
	if rds == nil {
		return false, errors.New("sign redis not configured")
	}

	key := "sign:nonce:" + accessKey + ":" + nonce
	return rds.SetnxExCtx(r.Context(), key, "1", signWindowSeconds)
}

func readSignBody(r *http.Request, maxBody int64) ([]byte, bool) {
	if r.Body == nil {
		return nil, true
	}

	body, err := io.ReadAll(io.LimitReader(r.Body, maxBody+1))
	if err != nil {
		return nil, false
	}
	r.Body = io.NopCloser(bytes.NewReader(body))
	return body, int64(len(body)) <= maxBody
}

func writeSignError(w http.ResponseWriter, r *http.Request, msg string) {
	httpx.OkJsonCtx(r.Context(), w, helper.Fail(errx.AuthErr(msg)))
}

func abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}

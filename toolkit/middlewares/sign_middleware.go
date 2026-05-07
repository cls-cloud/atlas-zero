package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func SignExecHandle(next http.HandlerFunc, rds *redis.Redis) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ---------- 解析 Header ----------
		accessKey := r.Header.Get("X-Key")
		timestampStr := r.Header.Get("X-Timestamp")
		nonce := r.Header.Get("X-Nonce")
		sign := r.Header.Get("X-Sign")
		if accessKey == "" || timestampStr == "" || nonce == "" || sign == "" {
			http.Error(w, "missing sign headers", http.StatusUnauthorized)
			return
		}
		// ---------- 时间戳校验（10 位秒级） ----------
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid timestamp", http.StatusUnauthorized)
			return
		}
		now := time.Now().Unix()
		if abs(now-timestamp) > 300 {
			http.Error(w, "timestamp expired", http.StatusUnauthorized)
			return
		}
		// ---------- accessKey -> secret ----------
		secret, err := getSecretByAccessKey(rds, accessKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// ---------- 构造 canonical string ----------
		canonical := buildCanonicalString(
			r.Method,
			r.URL.Path,
			timestamp,
			nonce, // 只参与签名
		)
		// ---------- 计算签名 ----------
		expectSign := hmacSign(canonical, secret)
		// ---------- 比较 ----------
		if !hmac.Equal([]byte(expectSign), []byte(sign)) {
			http.Error(w, "signature mismatch", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func getSecretByAccessKey(rds *redis.Redis, accessKey string) (string, error) {
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

/* ======================== Sign ======================== */
func buildCanonicalString(
	method, path string,
	timestamp int64,
	nonce string,
) string {
	return strings.Join([]string{
		method,
		path,
		strconv.FormatInt(timestamp, 10),
		nonce,
	}, "&")
}

func hmacSign(canonical, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(canonical))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}

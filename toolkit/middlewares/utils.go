package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
)

const EmptyBodySHA256 = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

type SignParams struct {
	Method     string
	Path       string
	Query      url.Values
	AccessKey  string
	Timestamp  int64
	Nonce      string
	BodySHA256 string
}

func (s SignParams) CanonicalString() string {
	bodyHash := s.BodySHA256
	if bodyHash == "" {
		bodyHash = EmptyBodySHA256
	}

	return strings.Join([]string{
		strings.ToUpper(s.Method),
		s.Path,
		s.Query.Encode(),
		s.AccessKey,
		strconv.FormatInt(s.Timestamp, 10),
		s.Nonce,
		bodyHash,
	}, "\n")
}

func Sign(canonical, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	_, _ = h.Write([]byte(canonical))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func SHA256Hex(body []byte) string {
	sum := sha256.Sum256(body)
	return hex.EncodeToString(sum[:])
}

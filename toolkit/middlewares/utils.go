package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type SignParams struct {
	Method    string
	Path      string
	ClientID  string
	Timestamp int64
	Nonce     string
}

func (s *SignParams) CanonicalString() string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%d\n%s",
		s.Method,
		s.Path,
		s.ClientID,
		s.Timestamp,
		s.Nonce,
	)
}

func Sign(canonical, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(canonical))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

package middlewares

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"ovra/toolkit/configshared"

	"github.com/stretchr/testify/require"
)

func TestApiEncryptMiddlewareDecryptsRequestAndEncryptsResponse(t *testing.T) {
	requestKey, err := rsa.GenerateKey(rand.Reader, 1024)
	require.NoError(t, err)
	responseKey, err := rsa.GenerateKey(rand.Reader, 1024)
	require.NoError(t, err)

	responsePublicDER, err := x509.MarshalPKIXPublicKey(&responseKey.PublicKey)
	require.NoError(t, err)

	cfg := configshared.ApiDecryptConfig{
		Enabled:    true,
		HeaderFlag: "encrypt-key",
		PrivateKey: base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(requestKey)),
		PublicKey:  base64.StdEncoding.EncodeToString(responsePublicDER),
	}

	secret := []byte("12345678901234567890123456789012")
	encryptedBody, err := aesECBEncrypt([]byte(`{"password":"s3cret"}`), secret)
	require.NoError(t, err)
	encryptedSecret, err := rsa.EncryptPKCS1v15(rand.Reader, &requestKey.PublicKey, []byte(base64.StdEncoding.EncodeToString(secret)))
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader([]byte(base64.StdEncoding.EncodeToString(encryptedBody))))
	req.Header.Set("encrypt-key", base64.StdEncoding.EncodeToString(encryptedSecret))
	rec := httptest.NewRecorder()

	middleware := ApiEncryptMiddleware(cfg)
	middleware(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.JSONEq(t, `{"password":"s3cret"}`, string(body))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"msg":"ok"}`))
	})(rec, req)

	responseEncryptKey := rec.Header().Get("encrypt-key")
	require.NotEmpty(t, responseEncryptKey)
	responseSecretCipher, err := base64.StdEncoding.DecodeString(responseEncryptKey)
	require.NoError(t, err)
	responseSecretBase64, err := rsa.DecryptPKCS1v15(rand.Reader, responseKey, responseSecretCipher)
	require.NoError(t, err)
	responseSecret, err := base64.StdEncoding.DecodeString(string(responseSecretBase64))
	require.NoError(t, err)
	require.Len(t, string(responseSecret), 32)

	responseCipher, err := base64.StdEncoding.DecodeString(rec.Body.String())
	require.NoError(t, err)
	responsePlain, err := aesECBDecrypt(responseCipher, responseSecret)
	require.NoError(t, err)
	require.JSONEq(t, `{"code":200,"msg":"ok"}`, string(responsePlain))
}

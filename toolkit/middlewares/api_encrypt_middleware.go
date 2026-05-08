package middlewares

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"ovra/toolkit/configshared"
	"ovra/toolkit/errx"
	"ovra/toolkit/helper"

	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	defaultEncryptHeader = "encrypt-key"
	apiEncryptMaxBody    = 10 << 20
)

func ApiEncryptMiddleware(c configshared.ApiDecryptConfig) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		if !c.Enabled {
			return next
		}

		header := strings.TrimSpace(c.HeaderFlag)
		if header == "" || strings.EqualFold(header, "X-Access-Token") {
			header = defaultEncryptHeader
		}

		return func(w http.ResponseWriter, r *http.Request) {
			encryptKey := strings.TrimSpace(r.Header.Get(header))
			if encryptKey == "" && !strings.EqualFold(header, defaultEncryptHeader) {
				encryptKey = strings.TrimSpace(r.Header.Get(defaultEncryptHeader))
			}
			if encryptKey == "" {
				next(w, r)
				return
			}

			secret, err := decryptSecret(encryptKey, c.PrivateKey)
			if err != nil {
				writeEncryptError(w, r, "请求加密密钥解密失败")
				return
			}

			if hasBody(r) {
				plainBody, err := decryptRequestBody(r, secret)
				if err != nil {
					writeEncryptError(w, r, "请求参数解密失败")
					return
				}
				r.Body = io.NopCloser(bytes.NewReader(plainBody))
				r.ContentLength = int64(len(plainBody))
				r.Header.Set("Content-Length", strconv.Itoa(len(plainBody)))
				r.Header.Set("Content-Type", "application/json")
			}

			recorder := newEncryptResponseWriter(w)
			next(recorder, r)

			if recorder.wroteHeader && recorder.statusCode >= http.StatusMultipleChoices {
				recorder.flushPlain()
				return
			}

			encryptedBody, responseEncryptKey, err := encryptResponseBody(recorder.body.Bytes(), c.PublicKey)
			if err != nil {
				writeEncryptError(w, r, "响应参数加密失败")
				return
			}

			headerMap := w.Header()
			addExposeHeader(headerMap, defaultEncryptHeader)
			headerMap.Set(defaultEncryptHeader, responseEncryptKey)
			headerMap.Set("Content-Type", "text/plain; charset=utf-8")
			headerMap.Set("Content-Length", strconv.Itoa(len(encryptedBody)))
			status := recorder.statusCode
			if status == 0 {
				status = http.StatusOK
			}
			w.WriteHeader(status)
			_, _ = w.Write([]byte(encryptedBody))
		}
	}
}

func decryptSecret(encryptKey string, privateKey string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptKey)
	if err != nil {
		return nil, err
	}
	key, err := parseRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	plain, err := rsa.DecryptPKCS1v15(rand.Reader, key, ciphertext)
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(string(plain))
}

func decryptRequestBody(r *http.Request, secret []byte) ([]byte, error) {
	body, err := io.ReadAll(io.LimitReader(r.Body, apiEncryptMaxBody+1))
	if err != nil {
		return nil, err
	}
	if int64(len(body)) > apiEncryptMaxBody {
		return nil, errors.New("request body too large")
	}
	bodyText := strings.TrimSpace(string(body))
	var quoted string
	if err := json.Unmarshal(body, &quoted); err == nil {
		bodyText = quoted
	}
	ciphertext, err := base64.StdEncoding.DecodeString(bodyText)
	if err != nil {
		return nil, err
	}
	return aesECBDecrypt(ciphertext, secret)
}

func encryptResponseBody(body []byte, publicKey string) (string, string, error) {
	secret, err := randomAESKey(32)
	if err != nil {
		return "", "", err
	}

	ciphertext, err := aesECBEncrypt(body, []byte(secret))
	if err != nil {
		return "", "", err
	}

	pub, err := parseRSAPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	keyCipher, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(base64.StdEncoding.EncodeToString([]byte(secret))))
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), base64.StdEncoding.EncodeToString(keyCipher), nil
}

func randomAESKey(length int) (string, error) {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	buf := make([]byte, length)
	max := big.NewInt(int64(len(letters)))
	for i := range buf {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		buf[i] = letters[n.Int64()]
	}
	return string(buf), nil
}

func aesECBEncrypt(plain []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plain = pkcs7Pad(plain, block.BlockSize())
	ciphertext := make([]byte, len(plain))
	for start := 0; start < len(plain); start += block.BlockSize() {
		block.Encrypt(ciphertext[start:start+block.BlockSize()], plain[start:start+block.BlockSize()])
	}
	return ciphertext, nil
}

func aesECBDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) == 0 || len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("invalid ciphertext size")
	}
	plain := make([]byte, len(ciphertext))
	for start := 0; start < len(ciphertext); start += block.BlockSize() {
		block.Decrypt(plain[start:start+block.BlockSize()], ciphertext[start:start+block.BlockSize()])
	}
	return pkcs7Unpad(plain, block.BlockSize())
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, errors.New("invalid pkcs7 data")
	}
	padding := int(data[len(data)-1])
	if padding == 0 || padding > blockSize || padding > len(data) {
		return nil, errors.New("invalid pkcs7 padding")
	}
	for _, v := range data[len(data)-padding:] {
		if int(v) != padding {
			return nil, errors.New("invalid pkcs7 padding")
		}
	}
	return data[:len(data)-padding], nil
}

func parseRSAPrivateKey(keyText string) (*rsa.PrivateKey, error) {
	der, err := decodeKey(keyText)
	if err != nil {
		return nil, err
	}
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	key, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not rsa private key")
	}
	return rsaKey, nil
}

func parseRSAPublicKey(keyText string) (*rsa.PublicKey, error) {
	der, err := decodeKey(keyText)
	if err != nil {
		return nil, err
	}
	if key, err := x509.ParsePKIXPublicKey(der); err == nil {
		rsaKey, ok := key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("not rsa public key")
		}
		return rsaKey, nil
	}
	return x509.ParsePKCS1PublicKey(der)
}

func decodeKey(keyText string) ([]byte, error) {
	keyText = strings.TrimSpace(keyText)
	if keyText == "" {
		return nil, errors.New("empty key")
	}

	raw, err := base64.StdEncoding.DecodeString(keyText)
	if err != nil {
		raw = []byte(keyText)
	}
	if block, _ := pem.Decode(raw); block != nil {
		return block.Bytes, nil
	}
	return raw, nil
}

type encryptResponseWriter struct {
	http.ResponseWriter
	body        *bytes.Buffer
	statusCode  int
	wroteHeader bool
}

func newEncryptResponseWriter(w http.ResponseWriter) *encryptResponseWriter {
	return &encryptResponseWriter{
		ResponseWriter: w,
		body:           bytes.NewBuffer(nil),
	}
}

func (w *encryptResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *encryptResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.wroteHeader = true
}

func (w *encryptResponseWriter) Write(data []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.body.Write(data)
}

func (w *encryptResponseWriter) flushPlain() {
	status := w.statusCode
	if status == 0 {
		status = http.StatusOK
	}
	w.ResponseWriter.WriteHeader(status)
	_, _ = w.ResponseWriter.Write(w.body.Bytes())
}

func hasBody(r *http.Request) bool {
	return r.Body != nil && r.Body != http.NoBody
}

func writeEncryptError(w http.ResponseWriter, r *http.Request, msg string) {
	httpx.OkJsonCtx(r.Context(), w, helper.Fail(errx.BizErr(msg)))
}

func addExposeHeader(header http.Header, name string) {
	const exposeHeader = "Access-Control-Expose-Headers"
	current := header.Get(exposeHeader)
	if current == "" {
		header.Set(exposeHeader, name)
		return
	}

	for _, item := range strings.Split(current, ",") {
		if strings.EqualFold(strings.TrimSpace(item), name) {
			return
		}
	}
	header.Set(exposeHeader, current+", "+name)
}

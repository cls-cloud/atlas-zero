package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type UserClaims struct {
	UserId   string `json:"userId"`
	TenantId string `json:"tenantId"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 token
func GenerateToken(userId, tenantId, secretKey string, expireInSeconds int64) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expireInSeconds) * time.Second)
	UserClaim := &UserClaims{
		UserId:   userId,
		TenantId: tenantId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken 解析 token
func AnalyseToken(tokenString, secretKey string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// GetUUID 生成唯一码
func GetUUID() string {
	return uuid.New().String()
}

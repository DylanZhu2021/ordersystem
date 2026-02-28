package utility

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   int64  `json:"userId"`
	UserType string `json:"userType"` // "user" or "admin"
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT token
func GenerateToken(ctx context.Context, userId int64, userType string) (string, error) {
	secret := g.Cfg().MustGet(ctx, "app.jwt.secret").String()
	expireKey := "app.jwt.expire"
	if userType == "admin" {
		expireKey = "app.jwt.adminExpire"
	}
	expire := g.Cfg().MustGet(ctx, expireKey).Int64()

	claims := Claims{
		UserId:   userId,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken 解析 JWT token
func ParseToken(ctx context.Context, tokenString string) (*Claims, error) {
	secret := g.Cfg().MustGet(ctx, "app.jwt.secret").String()
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

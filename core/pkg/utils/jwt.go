package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("tiktok*&2023%%125KBS627$*^")

type Claims struct {
	UserId int64 `json:"id"`
	jwt.StandardClaims
}

// GenerateToken 签发用户token
func GenerateToken(id int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tiktok",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

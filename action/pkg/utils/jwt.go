package utils

import (
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("tiktok*&2023%%125KBS627$*^")

type Claims struct {
	UserId int64 `json:"id"`
	jwt.StandardClaims
}

// GetUserId 获取用户id
func GetUserId(token string) int64 {
	tokenClaims, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims.UserId
		}
	}
	return 0
}

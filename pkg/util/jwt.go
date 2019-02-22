package util

import (
	"github.com/TopJohn/go-web-framework/pkg/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(conf.AppConfig.JwtSecret)

type Claims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func GenerateToken(userId int) (string, error) {
	now := time.Now()
	expireTime := now.Add(2 * time.Hour)
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

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

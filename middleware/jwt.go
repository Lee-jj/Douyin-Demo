package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("this is a jwt-go key")

type Claims struct {
	UserID   uint `json:"user_id"`
	userName string
	jwt.StandardClaims
}

func GreateToken(userID uint, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := time.Now().Add(time.Hour * 24)
	claims := Claims{
		UserID:   userID,
		userName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    "leechee",
			Subject:   "douyinDemo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

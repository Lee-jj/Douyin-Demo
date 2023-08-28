package middleware

import (
	"DOUYIN-DEMO/common"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SecretKey = []byte("this is a jwt-go key")

type Claims struct {
	UserID   int64 `json:"user_id"`
	userName string
	jwt.StandardClaims
}

func GreateToken(userID int64, userName string) (string, error) {
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

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware jwt check the token.")

		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1,
				StatusMsg:  common.ErrorHasNoToken.Error(),
			})
			c.Abort()
			return
		}

		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1,
				StatusMsg:  common.ErrorTokenFaild.Error(),
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1,
				StatusMsg:  common.ErrorTokenExpired.Error(),
			})
			c.Abort()
			return
		}

		// get user_id(host_id) and username
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.userName)

		c.Next()
	}
}

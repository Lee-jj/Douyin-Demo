package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegisterPesponse struct {
	common.Response
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	tokenResponse, err := service.UserRegisterService(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterPesponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, UserRegisterPesponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserID: tokenResponse.UserID,
		Token:  tokenResponse.Token,
	})
}

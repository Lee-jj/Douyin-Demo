package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	common.Response
	service.TokenResponse
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userTokenLoginResponse, err := service.UserLoginService(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		TokenResponse: userTokenLoginResponse,
	})
}

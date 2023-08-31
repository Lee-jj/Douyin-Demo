package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	common.Response
	service.UserInfoResponse `json:"user"`
}

type UserRegisterPesponse struct {
	common.Response
	service.TokenResponse
}

type UserLoginResponse struct {
	common.Response
	service.TokenResponse
}

/*************** User Infomatino Module ***************/
func UserInfo(c *gin.Context) {
	guestID := c.Query("user_id")
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)

	userInfoObjectResponse, err := service.UserInfoService(hostID, guestID)
	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		Response:         common.Response{StatusCode: 0},
		UserInfoResponse: userInfoObjectResponse,
	})
}

/*************** User Register Module ***************/
func UserRegister(c *gin.Context) {
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
			StatusMsg:  "注册&登录成功",
		},
		TokenResponse: tokenResponse,
	})
}

/*************** User Login Module ***************/
func UserLogin(c *gin.Context) {
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

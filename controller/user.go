package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/model"
	"DOUYIN-DEMO/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		// Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		// IsFollow:      true,
	},
}

type UserInfoResponse struct {
	common.Response
	service.UserInfoObjectResponse `json:"user"`
}

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
		Response:               common.Response{StatusCode: 0},
		UserInfoObjectResponse: userInfoObjectResponse,
	})
}

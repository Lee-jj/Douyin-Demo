package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	common.Response
	UserList []service.UserInfoResponse `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	toUserID := c.Query("to_user_id")
	actionType := c.Query("action_type")

	if hostID == toUserID {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "自己不能关注自己",
		})
		return
	}

	err := service.RelationActionService(hostID, toUserID, actionType)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "关注/取消关注 成功",
	})
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	UserID := c.Query("user_id")

	userList, _ := service.RelationFollowListService(hostID, UserID)

	if len(userList) == 0 {
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "用户关注列表为空",
			},
			UserList: nil,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "get follow list success.",
			},
			UserList: userList,
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	UserID := c.Query("user_id")

	userList, _ := service.RelationFollowerListService(hostID, UserID)

	if len(userList) == 0 {
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "用户粉丝列表为空",
			},
			UserList: nil,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "get follower list success.",
			},
			UserList: userList,
		})
	}
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {

}

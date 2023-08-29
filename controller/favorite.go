package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	videoID := c.Query("video_id")
	actionType := c.Query("action_type")

	err := service.FavoriteActionService(hostID, videoID, actionType)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "点赞/取消点赞 成功",
	})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}

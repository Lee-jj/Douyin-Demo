package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	common.Response
	VideoList []service.FeedVideoResponse `json:"video_list,omitempty"`
	NextTime  int64                       `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	strToken := c.Query("token")
	strLatestTime := c.Query("latest_time")

	latestTime, err := strconv.ParseInt(strLatestTime, 10, 32)
	if err != nil {
		latestTime = 0
	}
	// fmt.Printf("token: %v; latestTime: %v.\n", strToken, latestTime)
	videoList, _ := service.GetFeed(latestTime)

	feedVideoResponse, nextTime := service.FeedService(strToken, videoList)

	if len(feedVideoResponse) == 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  common.Response{StatusCode: 0, StatusMsg: "视频库为空"},
			VideoList: nil,
			NextTime:  0,
		})
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: feedVideoResponse,
		NextTime:  nextTime,
	})
}

package controller

import (
	"DOUYIN-DEMO/common"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	common.Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	strToken := c.Query("token")

	strLatestTime := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(strLatestTime, 10, 32)
	if err != nil {
		latestTime = 0
	}
	fmt.Printf("token: %v; latestTime: %v.\n", strToken, latestTime)

	// videoList, _ :=

	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}

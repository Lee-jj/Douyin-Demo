package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	common.Response
	VideoList []service.FeedVideoResponse `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "读取文件失败",
		})
		return
	}

	userID, videoName, err := service.GetPlayURL(token, title, file)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	videoPath := filepath.Join("./public", videoName)
	if err = c.SaveUploadedFile(file, videoPath); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "视频上传失败",
		})
		return
	}

	coverName := strings.Replace(videoName, ".mp4", ".jpeg", 1)
	err = service.GetCoverURL(videoName, coverName, 1)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	playURL := "http://192.168.31.246:8080/static/" + videoName
	coverURL := "http://192.168.31.246:8080/static/" + coverName
	err = service.CreateVideo(userID, playURL, coverURL, title)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  videoName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	guestID := c.Query("user_id")

	// 目标：获得guest_id的所有投稿视频
	feedVideoResponse, err := service.PublishListService(token, guestID)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:  common.Response{StatusCode: 1, StatusMsg: err.Error()},
			VideoList: nil,
		})
		return
	}

	if len(feedVideoResponse) == 0 {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "视频库为空",
			},
			VideoList: nil,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			VideoList: feedVideoResponse,
		})
	}

}

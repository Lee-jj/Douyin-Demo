package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	common.Response
	VideoList []service.VideoResponse `json:"video_list"`
}

/*************** Publish Action Module ***************/
// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	title := c.PostForm("title")
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "读取文件失败",
		})
		return
	}

	videoPath := filepath.Join("./public", filepath.Base(file.Filename))
	if err = c.SaveUploadedFile(file, videoPath); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "视频上传失败",
		})
		return
	}

	err = service.PublishService(hostID, videoPath, title)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

/*************** Publish List Module ***************/
// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	guestID := c.Query("user_id")
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)

	feedVideoResponse, err := service.PublishListService(hostID, guestID)
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

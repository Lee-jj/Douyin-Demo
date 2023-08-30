package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	common.Response
	CommentList []service.CommentResponse `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	common.Response
	Comment service.CommentResponse `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	actionType := c.Query("action_type")
	videoID := c.Query("video_id")
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	commentText := c.Query("comment_text")
	commentID := c.Query("comment_id")

	if actionType == "1" {
		// create comment
		commentResponse, err := service.CommentActionServiceCreate(hostID, videoID, commentText)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  "创建评论失败",
				},
			})
			return
		}

		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "创建评论成功",
			},
			Comment: commentResponse,
		})

	} else if actionType == "2" {
		// delete comment
		err := service.CommentActionServiceDelete(videoID, commentID)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  "删除评论失败",
				},
			})
			return
		}

		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "删除评论成功",
			},
		})

	} else {
		// error
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "非法评论操作",
			},
		})
		return
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	videoID := c.Query("video_id")

	commentList, _ := service.CommentListService(hostID, videoID)

	if len(commentList) == 0 {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "评论列表为空",
			},
			CommentList: nil,
		})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "get comment list success",
			},
			CommentList: commentList,
		})
	}
}

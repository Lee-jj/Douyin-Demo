package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/service"

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
	token := c.Query("token")
	actionType := c.Query("action_type")

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

}

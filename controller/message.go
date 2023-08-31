package controller

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/model"
	"DOUYIN-DEMO/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatResponse struct {
	common.Response
	MessageList []model.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	geustID := c.Query("to_user_id")
	actionType := c.Query("action_type")
	content := c.Query("content")

	if actionType == "1" {
		err := service.MessageActionService(hostID, geustID, content)
		if err != nil {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  "message action success.",
		})

	} else {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  common.ErrorWrongArgument.Error(),
		})
		return
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	hostIDAny, _ := c.Get("host_id")
	hostID := strconv.FormatInt(hostIDAny.(int64), 10)
	geustID := c.Query("to_user_id")
	// preMsgTime := c.Query("pre_msg_time")

	messageList, _ := service.MessageListService(hostID, geustID)

	if len(messageList) == 0 {
		c.JSON(http.StatusOK, ChatResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "聊天记录为空",
			},
			MessageList: nil,
		})
	} else {
		c.JSON(http.StatusOK, ChatResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "get chat list success.",
			},
			MessageList: messageList,
		})
	}
}

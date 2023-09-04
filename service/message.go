package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"strconv"
	"time"
)

func MessageActionService(fromUserID, toUserID, content string) error {
	fromUserIDInt, _ := strconv.ParseInt(fromUserID, 10, 64)
	toUserIDInt, _ := strconv.ParseInt(toUserID, 10, 64)

	tempMessage := model.Message{
		ToUserID:   toUserIDInt,
		FromUserID: fromUserIDInt,
		Content:    content,
		CreateAt:   time.Now().Unix(),
	}
	err := dao.CreateMessage(&tempMessage)

	return err
}

func MessageListService(fromUserID, toUserID, preMsgTime string) ([]model.Message, error) {
	fromUserIDInt, _ := strconv.ParseInt(fromUserID, 10, 64)
	toUserIDInt, _ := strconv.ParseInt(toUserID, 10, 64)
	preMsgTimeInt, _ := strconv.ParseInt(preMsgTime, 10, 64)

	var messageList []model.Message
	err := dao.GetMessageList(fromUserIDInt, toUserIDInt, preMsgTimeInt, &messageList)
	if err != nil {
		return []model.Message{}, nil
	}

	return messageList, nil
}

package dao

import "DOUYIN-DEMO/model"

func CreateMessage(tempMessage *model.Message) error {
	err := DB.Model(&model.Message{}).Create(tempMessage).Error
	return err
}

func GetMessageList(fromUserID, toUserID, preMsgTime int64, messageList *[]model.Message) error {
	err := DB.Model(&model.Message{}).Where("(from_user_id = ? AND to_user_id = ? AND create_at > ?) OR (from_user_id = ? AND to_user_id = ? AND create_at > ?)", fromUserID, toUserID, preMsgTime, toUserID, fromUserID, preMsgTime).Order("create_at").First(messageList).Error
	return err
}

func GetNewestMessage(fromUserID, toUserID int64, message *model.Message) error {
	err := DB.Model(&model.Message{}).Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", fromUserID, toUserID, toUserID, fromUserID).Order("create_at desc").First(message).Error
	return err
}

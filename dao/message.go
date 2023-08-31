package dao

import "DOUYIN-DEMO/model"

func CreateMessage(tempMessage *model.Message) error {
	err := DB.Model(&model.Message{}).Create(tempMessage).Error
	return err
}

func GetMessageList(fromUserID, toUserID int64, messageList *[]model.Message) error {
	err := DB.Model(&model.Message{}).Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", fromUserID, toUserID, toUserID, fromUserID).Order("create_at").Find(messageList).Error
	return err
}

func GetNewestMessage(fromUserID, toUserID int64, message *model.Message) error {
	err := DB.Model(&model.Message{}).Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", fromUserID, toUserID, toUserID, fromUserID).Order("create_at").First(message).Error
	return err
}

package dao

import "DOUYIN-DEMO/model"

func CreateComment(tempComment *model.Comment) error {
	err := DB.Model(&model.Comment{}).Create(tempComment).Error
	return err
}

func DeleteComment(tempComment *model.Comment) error {
	err := DB.Model(&model.Comment{}).Delete(tempComment).Error
	return err
}

func GetCommentByVideoID(videoID int64, tempCommentList *[]model.Comment) error {
	err := DB.Model(&model.Comment{}).Where("video_id = ?", videoID).Find(tempCommentList).Error
	return err
}

package dao

import "DOUYIN-DEMO/model"

func GetVideoByTime(timeFormat string, videoNum int, videoList *[]model.Video) error {
	err := DB.Model(&model.Video{}).Where("created_at < ?", timeFormat).Order("created_at desc").Limit(videoNum).First(videoList).Error
	return err
}

func GetVideoByUserID(guestID uint, videoList *[]model.Video) error {
	err := DB.Model(&model.Video{}).Where("auther_id = ?", guestID).Find(videoList).Error
	return err
}

package dao

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/model"
)

func GetVideoByTime(timeFormat string, videoNum int, videoList *[]model.Video) error {
	err := DB.Model(&model.Video{}).Where("created_at < ?", timeFormat).Order("created_at desc").Limit(videoNum).First(videoList).Error
	return err
}

func GetVideoByUserID(guestID int64, videoList *[]model.Video) error {
	err := DB.Model(&model.Video{}).Where("author_id = ?", guestID).Find(videoList).Error
	return err
}

func CreateVideo(tempVideo *model.Video) error {
	err := DB.Model(&model.Video{}).Create(&tempVideo).Error
	if err != nil {
		return common.ErrorCreateVideoFaild
	}
	return nil
}

func GetVideoNumByUserID(guestID int64, videoNum *int64) error {
	err := DB.Model(&model.Video{}).Where("author_id = ?", guestID).Count(videoNum).Error
	if err != nil {
		return err
	}
	return nil
}

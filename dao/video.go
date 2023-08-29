package dao

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/model"

	"gorm.io/gorm"
)

func GetVideoByTime(timeFormat string, videoNum int, videoList *[]model.Video) error {
	err := DB.Model(&model.Video{}).Where("created_at < ?", timeFormat).Order("created_at desc").Limit(videoNum).First(videoList).Error
	return err
}

func GetVideoByUserID(guestID int64, videoList *[]model.Video) error {
	err := DB.Model(&model.Video{}).Where("author_id = ?", guestID).Find(videoList).Error
	return err
}

func GetVideoByVideoID(videoID int64, video *model.Video) error {
	err := DB.Model(&model.Video{}).Where("video_id = ?", videoID).First(video).Error
	return err
}

func CreateVideo(tempVideo *model.Video) error {
	err := DB.Model(&model.Video{}).Create(&tempVideo).Error
	if err != nil {
		return common.ErrorCreateVideoFaild
	}
	return nil
}

func GetFavoriteVideoByUserID(guestID int64, favoriteList *[]model.Favorite) error {
	err := DB.Model(&model.Favorite{}).Where("user_id = ?", guestID).Find(favoriteList).Error
	return err
}

func AddVideoFavoriteCount(videoID, addNum int64) error {
	err := DB.Model(&model.Video{}).Where("video_id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count + ?", addNum)).Error
	return err
}

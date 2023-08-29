package dao

import "DOUYIN-DEMO/model"

func GetFavorite(userID, videoID int64, tempFavorite *model.Favorite) error {
	err := DB.Model(&model.Favorite{}).Where("user_id = ?", userID).Where("video_id = ?", videoID).First(tempFavorite).Error
	return err
}

func CreateFavorite(tempFavorite *model.Favorite) error {
	err := DB.Model(&model.Favorite{}).Create(tempFavorite).Error
	return err
}

func DeleteFavorite(tempFavorite *model.Favorite) error {
	err := DB.Model(&model.Favorite{}).Delete(tempFavorite).Error
	return err
}

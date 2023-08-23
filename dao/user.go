package dao

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/model"
	"errors"

	"gorm.io/gorm"
)

func GetUserByID(userID uint, tempUser *model.User) error {
	err := DB.Where("id=?", userID).First(tempUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// this user not in database
		return common.ErrorUserNotFound
	}

	if err != nil {
		return common.ErrorSQLFaild
	}

	// this user in database
	return nil
}

func GetUserByName(username string, tempUser *model.User) error {
	err := DB.Where("name=?", username).First(tempUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// this user not in database
		return common.ErrorUserNotFound
	}

	if err != nil {
		return common.ErrorSQLFaild
	}

	// this user in database
	return nil
}

func CreateUser(tempUser *model.User) error {
	err := DB.Model(&model.User{}).Create(&tempUser).Error
	if err != nil {
		return common.ErrorCreateUserFaild
	}

	return nil
}

package dao

import "DOUYIN-DEMO/model"

func GetRelation(hostID, toUserID int64, tempRelation *model.Relation) error {
	err := DB.Model(&model.Relation{}).Where("host_id = ?", hostID).Where("to_user_id = ?", toUserID).First(tempRelation).Error
	return err
}

func CreateRelation(tempRelation *model.Relation) error {
	err := DB.Model(&model.Relation{}).Create(tempRelation).Error
	return err
}

func DeleteRelation(tempRelation *model.Relation) error {
	err := DB.Model(&model.Relation{}).Delete(tempRelation).Error
	return err
}

func GetFollowList(userID int64, tempRelationList *[]model.Relation) error {
	err := DB.Model(&model.Relation{}).Where("host_id = ?", userID).Find(tempRelationList).Error
	return err
}

func GetFollowerList(userID int64, tempRelaionList *[]model.Relation) error {
	err := DB.Model(&model.Relation{}).Where("to_user_id = ?", userID).Find(tempRelaionList).Error
	return err
}

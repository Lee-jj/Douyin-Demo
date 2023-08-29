package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"strconv"
)

func FavoriteActionService(hostID, videoID, actionType string) error {
	hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)
	videoIDInt, _ := strconv.ParseInt(videoID, 10, 64)

	tempFavorite := model.Favorite{
		UserID:  hostIDInt,
		VideoID: videoIDInt,
	}
	err := dao.GetFavorite(hostIDInt, videoIDInt, &tempFavorite)

	if actionType == "1" {
		// like
		if err == nil {
			// already liked
			return common.ErrorAlreadyLiked
		} else {
			_ = dao.CreateFavorite(&tempFavorite)
		}
	} else if actionType == "2" {
		// cancel like
		if err != nil {
			// not liked
			return common.ErrorNotLiked
		} else {
			_ = dao.DeleteFavorite(&tempFavorite)
		}
	} else {
		return common.ErrorWrongArgument
	}

	return nil
}

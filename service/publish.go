package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"strconv"
)

func PublishListService(token, guestID string) ([]FeedVideoResponse, error) {
	var hasToken bool
	if token == "" {
		hasToken = false
	} else {
		hasToken = true
	}

	guestIDInt, err := strconv.ParseUint(guestID, 10, 64)
	if err != nil {
		return nil, err
	}

	tempUser := model.User{}
	err = dao.GetUserByID(uint(guestIDInt), &tempUser)
	if err != nil {
		return nil, err
	}

	feedUserInfo := FeedUserInfo{
		ID:             tempUser.ID,
		Name:           tempUser.Name,
		FollowCount:    tempUser.FollowCount,
		FollowerCount:  tempUser.FollowerCount,
		Avatar:         tempUser.Avatar,
		Backgroundmage: tempUser.BackgroundImage,
		TotalFavorite:  tempUser.TotalFavorited,
		FavoriteCount:  tempUser.FavoriteCount,
		IsFollow:       false,
	}
	if hasToken {
		feedUserInfo.IsFollow = 
	}

	videoList := []model.Video{}
	err = dao.GetVideoByUserID(uint(guestIDInt), &videoList)

	feedVideoResponse := []FeedVideoResponse{}
	for _, video := range videoList {

	}
}

package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
	"strconv"
	"time"
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
		tokenClaims, err1 := middleware.ParseToken(token)
		if err1 == nil && tokenClaims.ExpiresAt >= time.Now().Unix() {
			feedUserInfo.IsFollow = IsFollow(tokenClaims.UserID, strconv.Itoa(int(tempUser.ID)))
		}
	}

	videoList := []model.Video{}
	feedVideoResponse := []FeedVideoResponse{}
	err = dao.GetVideoByUserID(uint(guestIDInt), &videoList)
	// the video list is null, it is not an error, so we return null []FeedVideoResponse{} and nil
	if err != nil {
		return feedVideoResponse, nil
	}

	for _, video := range videoList {
		tempVideo := FeedVideoResponse{}

		tempVideo.ID = video.ID
		tempVideo.Author = feedUserInfo
		tempVideo.PlayUrl = video.PlayUrl
		tempVideo.CoverUrl = video.CoverUrl
		tempVideo.FavoriteCount = video.FavoriteCount
		tempVideo.CommentCount = video.CommentCount
		tempVideo.Title = video.Title
		tempVideo.IsFavorite = false
		if hasToken {
			tokenClaims, err2 := middleware.ParseToken(token)
			if err2 == nil && tokenClaims.ExpiresAt >= time.Now().Unix() {
				// For now, let's assume that the host user doesn't like any video
				tempVideo.IsFavorite = false
			}
		}

		feedVideoResponse = append(feedVideoResponse, tempVideo)
	}

	return feedVideoResponse, nil
}

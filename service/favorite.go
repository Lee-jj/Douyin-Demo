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

func FavoriteListService(hostID, guestID string) ([]VideoResponse, error) {
	guestIDInt, _ := strconv.ParseInt(guestID, 10, 64)

	var tempFavorite []model.Favorite
	err := dao.GetFavoriteVideoByUserID(guestIDInt, &tempFavorite)
	if err != nil {
		return []VideoResponse{}, nil
	}

	var videoListResponse []VideoResponse
	for _, favorite := range tempFavorite {
		var tempVideoResponse VideoResponse

		videoID := favorite.VideoID
		var video model.Video
		err := dao.GetVideoByVideoID(videoID, &video)
		if err != nil {
			continue
		}

		tempVideoResponse.ID = video.ID
		tempVideoResponse.PlayUrl = video.PlayUrl
		tempVideoResponse.CoverUrl = video.CoverUrl
		tempVideoResponse.FavoriteCount = video.FavoriteCount
		tempVideoResponse.IsFavorite = true
		tempVideoResponse.Title = video.Title

		var tempUser UserInfoResponse
		var user model.User
		err = dao.GetUserByID(video.AuthorID, &user)
		if err != nil {
			continue
		}

		tempUser.UserID = user.ID
		tempUser.UserName = user.Name
		tempUser.FollowCount = user.FollowCount
		tempUser.FollowerCount = user.FollowerCount
		tempUser.Avatar = user.Avatar
		tempUser.BackgroundImage = user.BackgroundImage
		tempUser.Signature = user.Signature
		tempUser.TotalFavorited = user.TotalFavorited
		tempUser.FavoriteCount = user.FavoriteCount

		var workCount int64
		_ = dao.GetVideoNumByUserID(guestIDInt, &workCount)
		tempUser.WorkCount = workCount

		tempUser.IsFollow = IsFollow(hostID, guestID)

		tempVideoResponse.Author = tempUser

		videoListResponse = append(videoListResponse, tempVideoResponse)
	}

	return videoListResponse, nil
}

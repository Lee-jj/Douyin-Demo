package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
	"strconv"
	"time"
)

type FeedUserInfo struct {
	ID             int64  `json:"user_id"`
	Name           string `json:"name"`
	FollowCount    int64  `json:"follow_count"`
	FollowerCount  int64  `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
	Avatar         string `json:"avater"`
	Backgroundmage string `json:"background_image"`
	TotalFavorite  int64  `json:"total_favorite"`
	FavoriteCount  int64  `json:"favorite_count"`
}

type FeedVideoResponse struct {
	ID            int64        `json:"id"`
	Author        FeedUserInfo `json:"author"`
	PlayUrl       string       `json:"play_url"`
	CoverUrl      string       `json:"cover_url"`
	FavoriteCount int64        `json:"favorite_count"`
	CommentCount  int64        `json:"comment_count"`
	IsFavorite    bool         `json:"is_favorite"`
	Title         string       `json:"title"`
}

func GetFeed(lastTime int64) ([]model.Video, error) {
	maxVideoNum := 30

	if lastTime == 0 {
		lastTime = time.Now().Unix()
	}
	timeFormat := time.Unix(lastTime, 0).Format("2006-01-02 15:04:05")

	var videoList []model.Video
	err := dao.GetVideoByTime(timeFormat, maxVideoNum, &videoList)

	return videoList, err
}

func FeedService(token string, videoList []model.Video) ([]FeedVideoResponse, int64) {
	var tokenIsNil bool
	if token == "" {
		tokenIsNil = true
	} else {
		tokenIsNil = false
	}

	feedVideoResponse := []FeedVideoResponse{}
	var nextTime int64
	for _, video := range videoList {
		tempFeedUser := FeedUserInfo{}
		tempVideo := FeedVideoResponse{}
		tempUser := model.User{}

		err := dao.GetUserByID(video.AuthorID, &tempUser)
		if err == nil {
			tempFeedUser.ID = tempUser.ID
			tempFeedUser.Name = tempUser.Name
			tempFeedUser.FollowCount = tempUser.FollowCount
			tempFeedUser.FollowerCount = tempUser.FollowerCount
			tempFeedUser.Avatar = tempUser.Avatar
			tempFeedUser.Backgroundmage = tempUser.BackgroundImage
			tempFeedUser.TotalFavorite = tempUser.TotalFavorited
			tempFeedUser.FavoriteCount = tempUser.FavoriteCount
			tempFeedUser.IsFollow = false

			if !tokenIsNil {
				tokenClaims, err1 := middleware.ParseToken(token)
				// token not expired
				if err1 == nil && time.Now().Unix() <= tokenClaims.ExpiresAt {
					tempFeedUser.IsFollow = IsFollow(tokenClaims.UserID, strconv.Itoa(int(tempUser.ID)))
				}
			}
		}

		tempVideo.ID = video.ID
		tempVideo.Author = tempFeedUser
		tempVideo.PlayUrl = video.PlayUrl
		tempVideo.CoverUrl = video.CoverUrl
		tempVideo.CommentCount = video.CommentCount
		tempVideo.FavoriteCount = video.FavoriteCount
		tempVideo.Title = video.Title
		tempVideo.IsFavorite = false

		if !tokenIsNil {
			tokenClaims, err1 := middleware.ParseToken(token)
			// token not expired
			if err1 == nil && time.Now().Unix() <= tokenClaims.ExpiresAt {
				// For now, let's assume that the host user doesn't like any video
				tempVideo.IsFavorite = false
			}
		}

		feedVideoResponse = append(feedVideoResponse, tempVideo)
		nextTime = video.CreatedAt.Unix()
	}
	return feedVideoResponse, nextTime
}

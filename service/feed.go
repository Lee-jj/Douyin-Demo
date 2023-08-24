package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
	"strconv"
	"time"
)

type FeedUserInfo struct {
	ID             uint   `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	FollowCount    uint   `json:"follow_count,omitempty"`
	FollowerCount  uint   `json:"follower_count,omitempty"`
	IsFollow       bool   `json:"is_follow,omitempty"`
	Avatar         string `json:"avater,omitempty"`
	Backgroundmage string `json:"background_image,omitempty"`
	TotalFavorite  uint   `json:"total_favorite,omitempty"`
	FavoriteCount  uint   `json:"favorite_count,omitempty"`
}

type FeedVideoResponse struct {
	ID            uint         `json:"id,omitempty"`
	Author        FeedUserInfo `json:"author,omitempty"`
	PlayUrl       string       `json:"play_url,omitempty"`
	CoverUrl      string       `json:"cover_url,omitempty"`
	FavoriteCount uint         `json:"favorite_count,omitempty"`
	CommentCount  uint         `json:"comment_count,omitempty"`
	IsFavorite    bool         `json:"is_favorite,omitempty"`
	Title         string       `json:"title,omitempty"`
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

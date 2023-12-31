package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
	"time"
)

type UserInfoResponse struct {
	UserID          int64  `json:"id"`
	UserName        string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type VideoResponse struct {
	ID            int64            `json:"id"`
	Author        UserInfoResponse `json:"author"`
	PlayUrl       string           `json:"play_url"`
	CoverUrl      string           `json:"cover_url"`
	FavoriteCount int64            `json:"favorite_count"`
	CommentCount  int64            `json:"comment_count"`
	IsFavorite    bool             `json:"is_favorite"`
	Title         string           `json:"title"`
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

func FeedService(token string, videoList []model.Video) ([]VideoResponse, int64) {
	var hasToken bool
	if token == "" {
		hasToken = false
	} else {
		hasToken = true
	}

	feedVideoResponse := []VideoResponse{}
	var nextTime int64
	for _, video := range videoList {
		tempFeedUser := UserInfoResponse{}
		tempVideo := VideoResponse{}
		tempUser := model.User{}

		err := dao.GetUserByID(video.AuthorID, &tempUser)
		if err == nil {
			// has user info
			tempFeedUser.UserID = tempUser.ID
			tempFeedUser.UserName = tempUser.Name
			tempFeedUser.FollowCount = tempUser.FollowCount
			tempFeedUser.FollowerCount = tempUser.FollowerCount
			tempFeedUser.Avatar = tempUser.Avatar
			tempFeedUser.BackgroundImage = tempUser.BackgroundImage
			tempFeedUser.TotalFavorited = tempUser.TotalFavorited
			tempFeedUser.WorkCount = tempUser.WorkCount
			tempFeedUser.FavoriteCount = tempUser.FavoriteCount
			tempFeedUser.IsFollow = false

			if hasToken {
				tokenClaims, err1 := middleware.ParseToken(token)
				// token not expired
				if err1 == nil && time.Now().Unix() <= tokenClaims.ExpiresAt {
					tempFeedUser.IsFollow = IsFollow(tokenClaims.UserID, tempUser.ID)
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

		if hasToken {
			tokenClaims, err1 := middleware.ParseToken(token)
			// token not expired
			if err1 == nil && time.Now().Unix() <= tokenClaims.ExpiresAt {
				// For now, let's assume that the host user doesn't like any video
				// tempVideo.IsFavorite = false
				var tempFavorite model.Favorite
				if err := dao.GetFavorite(tokenClaims.UserID, video.ID, &tempFavorite); err == nil {
					tempVideo.IsFavorite = true
				}
			}
		}

		feedVideoResponse = append(feedVideoResponse, tempVideo)
		nextTime = video.CreatedAt.Unix()
	}
	return feedVideoResponse, nextTime
}

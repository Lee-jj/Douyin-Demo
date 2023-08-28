package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"strconv"
)

type UserInfoObjectResponse struct {
	UserID          int64  `json:"id"`
	UserName        string `json:"name"`
	FollowCount     int64  `json:"follow_count"`   // 关注数
	FollowerCount   int64  `json:"follower_count"` // 被关注数
	IsFollowe       bool   `json:"is_follow"`
	Avatar          string `json:"avatar"` // 头像
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"` // 总被点赞数
	FavoriteCount   int64  `json:"favorite_count"`  // 点赞数
}

func UserInfoService(guestID string) (UserInfoObjectResponse, error) {
	userInfoObjectResponse := UserInfoObjectResponse{}

	guestIDInt, err := strconv.ParseInt(guestID, 10, 64)
	if err != nil {
		return userInfoObjectResponse, err
	}

	var user model.User
	err = dao.GetUserByID(guestIDInt, &user)
	if err != nil {
		return userInfoObjectResponse, err
	}

	userInfoObjectResponse = UserInfoObjectResponse{
		UserID:          user.ID,
		UserName:        user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollowe:       false,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		TotalFavorited:  user.TotalFavorited,
		FavoriteCount:   user.FavoriteCount,
	}
	return userInfoObjectResponse, nil
}

func IsFollow(hostID int64, guestID string) bool {
	guestIDInt, err := strconv.ParseInt(guestID, 10, 64)
	if err != nil {
		return false
	}

	// For now, let's assume that the host user follows all users except himself
	return hostID != guestIDInt
}

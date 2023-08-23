package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"strconv"
)

type UserInfoObjectResponse struct {
	UserID          uint   `json:"user_id"`
	UserName        string `json:"user_name"`
	FollowCount     uint   `json:"follow_count"`   // 关注数
	FollowerCount   uint   `json:"follower_count"` // 被关注数
	IsFollowe       bool   `json:"is_follow"`
	Avatar          string `json:"avatar"` // 头像
	BackgroundImage string `json:"background_image"`
	TotalFavorited  uint   `json:"total_favorited"` // 总被点赞数
	FavoriteCount   uint   `json:"favorite_count"`  // 点赞数
}

func UserInfoService(guestID string) (UserInfoObjectResponse, error) {
	userInfoObjectResponse := UserInfoObjectResponse{}

	guestIDInt, err := strconv.ParseUint(guestID, 10, 64)
	if err != nil {
		return userInfoObjectResponse, err
	}

	var user model.User
	err = dao.GetUserByID(uint(guestIDInt), &user)
	if err != nil {
		return userInfoObjectResponse, err
	}

	userInfoObjectResponse = UserInfoObjectResponse{
		UserID:          user.Model.ID,
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

func IsFollow(hostID uint, guestID string) bool {
	guestIDInt, err := strconv.ParseUint(guestID, 10, 64)
	if err != nil {
		return false
	}

	// For now, let's assume that the host user follows all users except himself
	return hostID != uint(guestIDInt)
}

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
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"` // 点赞数
}

func UserInfoService(hostID, guestID string) (UserInfoObjectResponse, error) {
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
	var workCount int64
	_ = dao.GetVideoNumByUserID(guestIDInt, &workCount)

	userInfoObjectResponse = UserInfoObjectResponse{
		UserID:          user.ID,
		UserName:        user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollowe:       false,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       workCount,
		FavoriteCount:   user.FavoriteCount,
	}
	userInfoObjectResponse.IsFollowe = IsFollow(hostID, guestID)

	return userInfoObjectResponse, nil
}

func IsFollow(hostID, guestID string) bool {
	// For now, let's assume that the host user follows all users except himself
	return hostID != guestID
}

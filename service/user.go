package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
	"fmt"
	"strconv"
)

const (
	MaxUsernameLen = 16
	MaxPasswordLen = 16
	MinPasswordLen = 6
)

// type UserInfoObjectResponse struct {
// 	UserID          int64  `json:"id"`
// 	UserName        string `json:"name"`
// 	FollowCount     int64  `json:"follow_count"`   // 关注数
// 	FollowerCount   int64  `json:"follower_count"` // 被关注数
// 	IsFollowe       bool   `json:"is_follow"`
// 	Avatar          string `json:"avatar"` // 头像
// 	BackgroundImage string `json:"background_image"`
// 	Signature       string `json:"signature"`
// 	TotalFavorited  int64  `json:"total_favorited"` // 总被点赞数
// 	WorkCount       int64  `json:"work_count"`
// 	FavoriteCount   int64  `json:"favorite_count"` // 点赞数
// }

type TokenResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func UserInfoService(hostID, guestID string) (UserInfoResponse, error) {
	userInfoObjectResponse := UserInfoResponse{}

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

	userInfoObjectResponse = UserInfoResponse{
		UserID:          user.ID,
		UserName:        user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        false,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       workCount,
		FavoriteCount:   user.FavoriteCount,
	}
	userInfoObjectResponse.IsFollow = IsFollow(hostID, guestID)

	return userInfoObjectResponse, nil
}

func IsFollow(hostID, guestID string) bool {
	// For now, let's assume that the host user follows all users except himself
	return hostID != guestID
}

func UserLoginService(username, password string) (TokenResponse, error) {
	var tokenResponse TokenResponse

	// username and password valid
	err := isUserValid(username, password)
	if err != nil {
		return tokenResponse, err
	}

	// user has not registered
	var login model.User
	err = dao.GetUserByName(username, &login)
	if err != nil {
		return tokenResponse, err
	}

	// password wrong
	if password != login.Password {
		return tokenResponse, common.ErrorPasswordWrong
	}

	token, err := middleware.GreateToken(login.ID, login.Name)
	if err != nil {
		return tokenResponse, err
	}

	tokenResponse = TokenResponse{
		UserID: login.ID,
		Token:  token,
	}
	return tokenResponse, nil
}

func UserRegisterService(userName, passWord string) (TokenResponse, error) {
	var tokenResponse TokenResponse

	err := isUserValid(userName, passWord)
	if err != nil {
		return tokenResponse, err
	}

	newUser, err := CreateRegisterUser(userName, passWord)
	if err != nil {
		return tokenResponse, err
	}

	token, err := middleware.GreateToken(newUser.ID, newUser.Name)
	if err != nil {
		return tokenResponse, err
	}

	tokenResponse = TokenResponse{
		UserID: newUser.ID,
		Token:  token,
	}
	return tokenResponse, nil
}

func isUserValid(userName, passWord string) error {
	if len(userName) == 0 {
		return common.ErrorUserNameEmpty
	}
	if len(userName) > MaxUsernameLen {
		return common.ErrorUserNameInvalid
	}

	if len(passWord) == 0 {
		return common.ErrorPasswordEmpty
	}
	if len(passWord) < MinPasswordLen || len(passWord) > MaxPasswordLen {
		return common.ErrorPasswordInvalid
	}

	return nil
}

func CreateRegisterUser(username, password string) (model.User, error) {
	newUser := model.User{
		Name:            username,
		Password:        password,
		Avatar:          "http://192.168.31.246:8080/static/defaultAvatar.jpg",
		BackgroundImage: "http://192.168.31.246:8080/static/defaultBackground.jpg",
		Signature:       "I have nothing to say",
	}

	var tempUser model.User
	err := dao.GetUserByName(username, &tempUser)
	if err == nil {
		return newUser, common.ErrorUserExist
	}

	err = dao.CreateUser(&newUser)
	if err != nil {
		return newUser, err
	}

	fmt.Printf("create a new user named %v\n", newUser.Name)
	return newUser, nil
}

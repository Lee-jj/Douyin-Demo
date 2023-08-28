package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
	"fmt"
)

const (
	MaxUsernameLen = 16
	MaxPasswordLen = 16
	MinPasswordLen = 6
)

type TokenResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
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

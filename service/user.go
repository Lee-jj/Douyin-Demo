package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
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

	userInfoObjectResponse = UserInfoResponse{
		UserID:          user.ID,
		UserName:        user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        false,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
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
	// Passwords obtained from the database are decrypted first
	if !PasswordIsRight(password, login.Password) {
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
	// Passwords are encrypted before they are placed in the database.
	hashPassword, err := Encryption(password)
	if err != nil {
		return model.User{}, err
	}

	newUser := model.User{
		Name:            username,
		Password:        hashPassword,
		Avatar:          "http://192.168.31.246:8080/static/defaultAvatar.jpg",
		BackgroundImage: "http://192.168.31.246:8080/static/defaultBackground.jpg",
		Signature:       "I have nothing to say",
	}

	var tempUser model.User
	err = dao.GetUserByName(username, &tempUser)
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

func Encryption(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", common.ErrorEncrypteFaild
	}

	return string(hashPassword), nil
}

func PasswordIsRight(password, hashPassword string) bool {
	// Decryption
	bytePwd := []byte(password)
	byteHash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	return err == nil
}

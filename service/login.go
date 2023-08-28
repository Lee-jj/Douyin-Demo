package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
)

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

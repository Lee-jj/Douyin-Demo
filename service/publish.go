package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/middleware"
	"DOUYIN-DEMO/model"
	"bytes"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"gorm.io/gorm"
)

func PublishListService(token, guestID string) ([]FeedVideoResponse, error) {
	var hasToken bool
	if token == "" {
		hasToken = false
	} else {
		hasToken = true
	}

	guestIDInt, err := strconv.ParseUint(guestID, 10, 64)
	if err != nil {
		return nil, err
	}

	tempUser := model.User{}
	err = dao.GetUserByID(uint(guestIDInt), &tempUser)
	if err != nil {
		return nil, err
	}

	feedUserInfo := FeedUserInfo{
		ID:             tempUser.ID,
		Name:           tempUser.Name,
		FollowCount:    tempUser.FollowCount,
		FollowerCount:  tempUser.FollowerCount,
		Avatar:         tempUser.Avatar,
		Backgroundmage: tempUser.BackgroundImage,
		TotalFavorite:  tempUser.TotalFavorited,
		FavoriteCount:  tempUser.FavoriteCount,
		IsFollow:       false,
	}
	if hasToken {
		tokenClaims, err1 := middleware.ParseToken(token)
		if err1 == nil && tokenClaims.ExpiresAt >= time.Now().Unix() {
			feedUserInfo.IsFollow = IsFollow(tokenClaims.UserID, strconv.Itoa(int(tempUser.ID)))
		}
	}

	videoList := []model.Video{}
	feedVideoResponse := []FeedVideoResponse{}
	err = dao.GetVideoByUserID(uint(guestIDInt), &videoList)
	// the video list is null, it is not an error, so we return null []FeedVideoResponse{} and nil
	if err != nil {
		return feedVideoResponse, nil
	}

	for _, video := range videoList {
		tempVideo := FeedVideoResponse{}

		tempVideo.ID = video.ID
		tempVideo.Author = feedUserInfo
		tempVideo.PlayUrl = video.PlayUrl
		tempVideo.CoverUrl = video.CoverUrl
		tempVideo.FavoriteCount = video.FavoriteCount
		tempVideo.CommentCount = video.CommentCount
		tempVideo.Title = video.Title
		tempVideo.IsFavorite = false
		if hasToken {
			tokenClaims, err2 := middleware.ParseToken(token)
			if err2 == nil && tokenClaims.ExpiresAt >= time.Now().Unix() {
				// For now, let's assume that the host user doesn't like any video
				tempVideo.IsFavorite = false
			}
		}

		feedVideoResponse = append(feedVideoResponse, tempVideo)
	}

	return feedVideoResponse, nil
}

func GetPlayURL(token, title string, file *multipart.FileHeader) (uint, string, error) {
	if token == "" {
		return 0, "", common.ErrorHasNoToken
	}

	if title == "" {
		return 0, "", common.ErrorHasNoTitle
	}

	tokenClaims, err := middleware.ParseToken(token)
	if err != nil {
		return 0, "", common.ErrorTokenFaild
	}
	userID := tokenClaims.UserID

	// video path
	originName := filepath.Base(file.Filename)
	fileName := fmt.Sprintf("%d_%d_%s", userID, time.Now().Unix(), originName)
	// filePath := filepath.Join("/static", fileName)
	return userID, fileName, nil

}

func GetCoverURL(videoName, imageName string, frameNum int) error {
	videoPath := filepath.Join("./public", videoName)
	imagePath := filepath.Join("./public", imageName)

	fmt.Printf("videoPath: %v", videoPath)

	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return err
	}

	err = imaging.Save(img, imagePath)
	if err != nil {
		return err
	}

	return nil
}

func CreateVideo(userID uint, playURL, coverURL, title string) error {
	tempVideo := model.Video{
		Model:         gorm.Model{},
		AuthorID:      userID,
		PlayUrl:       playURL,
		CoverUrl:      coverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}
	err := dao.CreateVideo(&tempVideo)
	if err != nil {
		return err
	}
	return nil
}

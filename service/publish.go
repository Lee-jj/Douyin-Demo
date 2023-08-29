package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
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
)

func PublishListService(hostID, guestID string) ([]VideoResponse, error) {
	guestIDInt, _ := strconv.ParseInt(guestID, 10, 64)
	// hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)

	tempUser := model.User{}
	err := dao.GetUserByID(guestIDInt, &tempUser)
	if err != nil {
		return nil, err
	}

	feedUserInfo := UserInfoResponse{
		UserID:          tempUser.ID,
		UserName:        tempUser.Name,
		FollowCount:     tempUser.FollowCount,
		FollowerCount:   tempUser.FollowerCount,
		Avatar:          tempUser.Avatar,
		BackgroundImage: tempUser.BackgroundImage,
		TotalFavorited:  tempUser.TotalFavorited,
		WorkCount:       tempUser.WorkCount,
		FavoriteCount:   tempUser.FavoriteCount,
		IsFollow:        false,
	}

	feedUserInfo.IsFollow = IsFollow(hostID, guestID)

	videoList := []model.Video{}
	feedVideoResponse := []VideoResponse{}
	err = dao.GetVideoByUserID(guestIDInt, &videoList)
	// the video list is null, it is not an error, so we return null []FeedVideoResponse{} and nil
	if err != nil {
		return feedVideoResponse, nil
	}

	for _, video := range videoList {
		tempVideo := VideoResponse{}

		tempVideo.ID = video.ID
		tempVideo.Author = feedUserInfo
		tempVideo.PlayUrl = video.PlayUrl
		tempVideo.CoverUrl = video.CoverUrl
		tempVideo.FavoriteCount = video.FavoriteCount
		tempVideo.CommentCount = video.CommentCount
		tempVideo.Title = video.Title
		// For now, let's assume that the host user doesn't like any video
		tempVideo.IsFavorite = false
		var tempFavorite model.Favorite
		if err := dao.GetFavorite(guestIDInt, video.ID, &tempFavorite); err == nil {
			tempVideo.IsFavorite = true
		}

		feedVideoResponse = append(feedVideoResponse, tempVideo)
	}

	return feedVideoResponse, nil
}

func GetPlayURL(hostID, title string, file *multipart.FileHeader) (string, error) {
	if title == "" {
		return "", common.ErrorHasNoTitle
	}

	// video path
	originName := filepath.Base(file.Filename)
	fileName := fmt.Sprintf("%s_%d_%s", hostID, time.Now().Unix(), originName)
	// filePath := filepath.Join("/static", fileName)
	return fileName, nil

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

func CreateVideo(hostID string, playURL, coverURL, title string) error {
	userID, err := strconv.ParseInt(hostID, 10, 64)
	if err != nil {
		return err
	}

	tempVideo := model.Video{
		AuthorID:      userID,
		PlayUrl:       playURL,
		CoverUrl:      coverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}
	err = dao.CreateVideo(&tempVideo)
	if err != nil {
		return err
	}
	err = dao.AddUserWorkCount(userID)
	if err != nil {
		return err
	}

	return nil
}

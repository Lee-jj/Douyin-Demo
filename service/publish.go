package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func PublishListService(hostID, guestID string) ([]VideoResponse, error) {
	guestIDInt, _ := strconv.ParseInt(guestID, 10, 64)
	hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)

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
		IsFollow:        IsFollow(hostIDInt, guestIDInt),
		Avatar:          tempUser.Avatar,
		BackgroundImage: tempUser.BackgroundImage,
		TotalFavorited:  tempUser.TotalFavorited,
		WorkCount:       tempUser.WorkCount,
		FavoriteCount:   tempUser.FavoriteCount,
	}

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
		tempVideo.IsFavorite = false
		var tempFavorite model.Favorite
		if err := dao.GetFavorite(guestIDInt, video.ID, &tempFavorite); err == nil {
			tempVideo.IsFavorite = true
		}

		feedVideoResponse = append(feedVideoResponse, tempVideo)
	}

	return feedVideoResponse, nil
}

func getCover(videoPath, imagePath string, frameNum int) error {
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

func createVideo(hostID string, playURL, coverURL, title string) error {
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

func PublishService(userID, videoPath, title string) error {
	client := dao.GetMinio()
	playURL, err := client.UpLoadFile("video", videoPath, userID)
	if err != nil {
		return err
	}

	imagePath := strings.Replace(videoPath, ".mp4", ".jpg", 1)
	err = getCover(videoPath, imagePath, 1)
	if err != nil {
		return err
	}

	coverURL, err := client.UpLoadFile("image", imagePath, userID)
	if err != nil {
		return err
	}

	err = createVideo(userID, playURL, coverURL, title)
	if err != nil {
		return err
	}

	// delete the temp video and image
	err = os.Remove(videoPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = os.Remove(imagePath)
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

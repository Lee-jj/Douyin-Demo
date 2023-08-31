package service

import (
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"strconv"
	"time"
)

type CommentResponse struct {
	ID       int64            `json:"id"`
	User     UserInfoResponse `json:"user"`
	Content  string           `json:"content"`
	CreateAt string           `json:"create_date"`
}

func CommentActionServiceCreate(hostID, videoID, commentText string) (CommentResponse, error) {
	hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)
	videoIDInt, _ := strconv.ParseInt(videoID, 10, 64)
	createAt := time.Now().Format("01-02")

	tempComment := model.Comment{
		UserID:   hostIDInt,
		VideoID:  videoIDInt,
		Content:  commentText,
		CreateAt: createAt,
	}
	err := dao.CreateComment(&tempComment)
	if err != nil {
		return CommentResponse{}, err
	}

	var tempUser model.User
	err = dao.GetUserByID(hostIDInt, &tempUser)
	if err != nil {
		return CommentResponse{}, err
	}

	userResponse := UserInfoResponse{
		UserID:          tempUser.ID,
		UserName:        tempUser.Name,
		FollowCount:     tempUser.FollowCount,
		FollowerCount:   tempUser.FollowerCount,
		IsFollow:        false, // 自己不能关注自己
		Avatar:          tempUser.Avatar,
		BackgroundImage: tempUser.BackgroundImage,
		Signature:       tempUser.Signature,
		TotalFavorited:  tempUser.TotalFavorited,
		WorkCount:       tempUser.WorkCount,
		FavoriteCount:   tempUser.FavoriteCount,
	}

	commentResponse := CommentResponse{
		ID:       tempComment.ID,
		User:     userResponse,
		Content:  commentText,
		CreateAt: createAt,
	}

	// update the comment_count of videos table
	_ = dao.AddVideoCommentCount(videoIDInt, 1)

	return commentResponse, nil
}

func CommentActionServiceDelete(videoID, commentID string) error {
	commentIDInt, _ := strconv.ParseInt(commentID, 10, 64)
	videoIDInt, _ := strconv.ParseInt(videoID, 10, 64)
	var tempComment model.Comment
	tempComment.ID = commentIDInt

	err := dao.DeleteComment(&tempComment)
	if err == nil {
		// delete success
		// update the comment_count of videos table
		_ = dao.AddVideoCommentCount(videoIDInt, -1)
	}

	return err
}

func CommentListService(hostID, videoID string) ([]CommentResponse, error) {
	hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)
	videoIDInt, _ := strconv.ParseInt(videoID, 10, 64)

	var commentList []model.Comment
	err := dao.GetCommentByVideoID(videoIDInt, &commentList)
	if err != nil {
		return []CommentResponse{}, nil
	}

	var commentResponseList []CommentResponse
	for _, comment := range commentList {
		var tempComment CommentResponse

		tempComment.ID = comment.ID
		tempComment.Content = comment.Content
		tempComment.CreateAt = comment.CreateAt

		var user model.User
		err := dao.GetUserByID(comment.UserID, &user)
		if err != nil {
			continue
		}

		tempUser := UserInfoResponse{
			UserID:          user.ID,
			UserName:        user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        IsFollow(hostIDInt, user.ID),
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}

		tempComment.User = tempUser

		commentResponseList = append(commentResponseList, tempComment)
	}

	return commentResponseList, nil
}

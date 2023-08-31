package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"strconv"
)

type FriendUser struct {
	UserInfoResponse
	Message string `json:"message,omitempty"`
	MsgType int64  `json:"msgType,omitempty"`
}

func RelationActionService(hostID, toUserID, actionType string) error {
	hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)
	toUserIDInt, _ := strconv.ParseInt(toUserID, 10, 64)

	tempRelation := model.Relation{
		HostID:   hostIDInt,
		ToUserID: toUserIDInt,
	}
	err := dao.GetRelation(hostIDInt, toUserIDInt, &tempRelation)

	if actionType == "1" {
		// follow
		if err == nil {
			// already follow
			return common.ErrorAlreadyFollowed
		} else {
			_ = dao.CreateRelation(&tempRelation)
			_ = dao.AddFollowCount(hostIDInt, 1)
			_ = dao.AddFollowerCount(toUserIDInt, 1)
		}
	} else if actionType == "2" {
		// cancel follow
		if err != nil {
			return common.ErrorNotFollowed
		} else {
			_ = dao.DeleteRelation(&tempRelation)
			_ = dao.AddFollowCount(hostIDInt, -1)
			_ = dao.AddFollowerCount(toUserIDInt, -1)
		}
	} else {
		// illegal
		return common.ErrorWrongArgument
	}

	return nil
}

func RelationFollowListService(hostID, geustID string) ([]UserInfoResponse, error) {
	// geustID as host_id in relations table
	hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)
	geustIDInt, _ := strconv.ParseInt(geustID, 10, 64)

	var tempRelationList []model.Relation
	err := dao.GetFollowList(geustIDInt, &tempRelationList)
	if err != nil {
		return []UserInfoResponse{}, nil
	}

	var userList []UserInfoResponse
	for _, relation := range tempRelationList {
		var user UserInfoResponse

		var tempUser model.User
		err := dao.GetUserByID(relation.ToUserID, &tempUser)
		if err != nil {
			continue
		}

		user.UserID = tempUser.ID
		user.UserName = tempUser.Name
		user.FollowCount = tempUser.FollowCount
		user.FollowerCount = tempUser.FollowerCount
		user.IsFollow = IsFollow(hostIDInt, tempUser.ID)
		user.Avatar = tempUser.Avatar
		user.BackgroundImage = tempUser.BackgroundImage
		user.Signature = tempUser.Signature
		user.TotalFavorited = tempUser.TotalFavorited
		user.WorkCount = tempUser.WorkCount
		user.FavoriteCount = tempUser.FavoriteCount

		userList = append(userList, user)
	}

	return userList, nil
}

func RelationFollowerListService(hostID, geustID string) ([]UserInfoResponse, error) {
	// geustID as to_user_id in relations table
	hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)
	geustIDInt, _ := strconv.ParseInt(geustID, 10, 64)

	var tempRelationList []model.Relation
	err := dao.GetFollowerList(geustIDInt, &tempRelationList)
	if err != nil {
		return []UserInfoResponse{}, nil
	}

	var userList []UserInfoResponse
	for _, relation := range tempRelationList {
		var user UserInfoResponse

		var tempUser model.User
		err := dao.GetUserByID(relation.HostID, &tempUser)
		if err != nil {
			continue
		}

		user.UserID = tempUser.ID
		user.UserName = tempUser.Name
		user.FollowCount = tempUser.FollowCount
		user.FollowerCount = tempUser.FollowerCount
		user.IsFollow = IsFollow(hostIDInt, tempUser.ID)
		user.Avatar = tempUser.Avatar
		user.BackgroundImage = tempUser.BackgroundImage
		user.Signature = tempUser.Signature
		user.TotalFavorited = tempUser.TotalFavorited
		user.WorkCount = tempUser.WorkCount
		user.FavoriteCount = tempUser.FavoriteCount

		userList = append(userList, user)
	}

	return userList, nil
}

func RelationFriendListService(hostID, geustID string) ([]FriendUser, error) {
	// You can only talk to your own fans
	hostIDInt, _ := strconv.ParseInt(hostID, 10, 64)

	userList, _ := RelationFollowerListService(hostID, geustID)

	var friendUserList []FriendUser
	for _, user := range userList {
		var tempFriendUser FriendUser

		tempFriendUser.UserInfoResponse = user

		var message model.Message
		err := dao.GetNewestMessage(hostIDInt, user.UserID, &message)
		if err == nil {
			tempFriendUser.Message = message.Content
			if message.FromUserID == hostIDInt {
				tempFriendUser.MsgType = 1
			} else if message.ToUserID == hostIDInt {
				tempFriendUser.MsgType = 0
			} else {
				continue
			}
		}

		friendUserList = append(friendUserList, tempFriendUser)
	}

	return friendUserList, nil
}

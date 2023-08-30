package service

import (
	"DOUYIN-DEMO/common"
	"DOUYIN-DEMO/dao"
	"DOUYIN-DEMO/model"
	"strconv"
)

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

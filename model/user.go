package model

import "gorm.io/gorm"

// default table name is "users"
type User struct {
	gorm.Model
	Name            string `json:"name"` // 要求用户名不重复
	Password        string `json:"password"`
	FollowCount     uint   `json:"follow_count"`   // 关注数
	FollowerCount   uint   `json:"follower_count"` // 被关注数
	Avatar          string `json:"avatar"`         // 头像
	BackgroundImage string `json:"background_image"`
	TotalFavorited  uint   `json:"total_favorited"` // 总被点赞数
	FavoriteCount   uint   `json:"favorite_count"`  // 点赞数
}

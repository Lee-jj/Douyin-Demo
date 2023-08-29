package model

// default table name is "users"
type User struct {
	ID              int64  `gorm:"column:user_id; primary_key"`
	Name            string `gorm:"column:name"` // 要求用户名不重复
	Password        string `gorm:"column:password"`
	FollowCount     int64  `gorm:"column:follow_count"`   // 关注数
	FollowerCount   int64  `gorm:"column:follower_count"` // 被关注数
	Avatar          string `gorm:"column:avatar"`         // 头像
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
	TotalFavorited  int64  `gorm:"column:total_favorited"` // 总被点赞数
	WorkCount       int64  `gorm:"column:work_count"`
	FavoriteCount   int64  `gorm:"column:favorite_count"` // 点赞数
}

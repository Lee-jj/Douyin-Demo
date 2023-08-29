package model

// default table name is "favorites"
type Comment struct {
	ID       int64  `gorm:"column:comment_id; primary_key"`
	UserID   int64  `gorm:"column:user_id"`
	VideoID  int64  `gorm:"column:video_id"`
	Content  string `gorm:"column:content"`
	CreateAt string `gorm:"column:create_date"`
}

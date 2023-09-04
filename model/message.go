package model

// default table name is "messages"
type Message struct {
	ID         int64  `gorm:"column:message_id; primary_key" json:"id"`
	ToUserID   int64  `gorm:"column:to_user_id" json:"to_user_id"`
	FromUserID int64  `gorm:"column:from_user_id" json:"from_user_id"`
	Content    string `gorm:"column:content" json:"content"`
	CreateAt   int64  `gorm:"column:create_at" json:"create_time"` // yyyy-MM-dd HH:MM:ss
}

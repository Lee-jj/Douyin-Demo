package model

import "time"

// default table name is "videos"
type Video struct {
	ID            int64     `gorm:"column:video_id; primary_key;"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	AuthorID      int64     `gorm:"column:author_id"`
	PlayUrl       string    `gorm:"column:play_url"`
	CoverUrl      string    `gorm:"column:cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count"`
	Title         string    `gorm:"column:title"`
}

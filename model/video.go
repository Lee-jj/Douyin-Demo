package model

import "gorm.io/gorm"

// default table name is "videos"
type Video struct {
	gorm.Model
	AuthorID      uint   `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	Title         string `json:"title"`
}

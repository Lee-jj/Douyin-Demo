package model

// default table name is "favorites"
// If the corresponding user_id and video_id records exist in the table, the user likes the video. Otherwise, the user does not like the video
type Favorite struct {
	FavoriteID int64 `gorm:"column:favorite_id; primary_key"`
	UserID     int64 `gorm:"column:user_id"`
	VideoID    int64 `gorm:"column:video_id"`
}

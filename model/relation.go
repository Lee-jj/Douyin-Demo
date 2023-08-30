package model

// default table name is "relations"
// If the corresponding host_id and to_user_id records exist in the table, the host follows the to_user. Otherwise, the host does not follow the to_user
type Relation struct {
	ID       int64 `gorm:"column:relation_id; primary_key"`
	HostID   int64 `gorm:"column:host_id"`
	ToUserID int64 `gorm:"column:to_user_id"`
}

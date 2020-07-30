package models

import "time"

//StoreMedia ...
type StoreMedia struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name" form:"name"`
	Type      string     `json:"type" gorm:"type"`
	Thumbnail string     `json:"thumbnail" gorm:"thumbnail"`
	UserID    *string    `json:"user_id,omitempty" form:"user_id,omitempty"`
	StoreID   *string    `json:"store_id,omitempty" form:"store_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

//TableName :
func (StoreMedia) TableName() string {
	return "store_medias"
}

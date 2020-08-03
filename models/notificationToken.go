package models

import "time"

//NotificationToken ...
type NotificationToken struct {
	ID        string     `json:"id" form:"id"`
	Value     string     `json:"value" form:"value" gorm:"value"`
	Token     string     `json:"token" form:"token" gorm:"token"`
	UserID    *string    `json:"user_id" form:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

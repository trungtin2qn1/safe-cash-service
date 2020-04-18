package models

import "time"

//NotificationToken ...
type NotificationToken struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	Value     string     `json:"value,omitempty" form:"value,omitempty"`
	Token     string     `json:"token,omitempty" form:"token,omitempty"`
	UserID    *string    `json:"user_id,omitempty" form:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

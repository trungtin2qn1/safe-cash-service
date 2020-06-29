package models

import "time"

//UserLoginLog ...
type UserLoginLog struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	UserID    *string    `json:"user_id" form:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

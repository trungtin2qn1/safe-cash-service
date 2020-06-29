package models

import "time"

//UnlockingLog ...
type UnlockingLog struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	Content   string     `json:"content" form:"content"`
	IsSuccess *bool      `json:"is_success" form:"is_success"`
	UserID    *string    `json:"user_id" form:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

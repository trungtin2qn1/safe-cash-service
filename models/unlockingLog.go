package models

import "time"

//UnlockingLog ...
type UnlockingLog struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	Content   string     `json:"content,omitempty" form:"content,omitempty"`
	IsSuccess *bool      `json:"is_success,omitempty" form:"is_success,omitempty"`
	UserID    *string    `json:"user_id,omitempty" form:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

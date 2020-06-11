package models

import "time"

//Notification ...
type Notification struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	Title     string     `json:"title,omitempty" form:"title,omitempty"`
	Body      string     `json:"body,omitempty" form:"body,omitempty"`
	IsRead    bool       `json:"is_read" form:"is_read"`
	UserID    *string    `json:"user_id,omitempty" form:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

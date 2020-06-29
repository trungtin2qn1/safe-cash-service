package models

import "time"

//Notification ...
type Notification struct {
	ID        string     `json:"id" form:"id"`
	Title     string     `json:"title" form:"title"`
	Body      string     `json:"body" form:"body"`
	IsRead    bool       `json:"is_read" form:"is_read"`
	UserID    *string    `json:"user_id" form:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

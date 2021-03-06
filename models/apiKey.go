package models

import "time"

//ApiKey ...
type ApiKey struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	Value     string     `json:"value,omitempty" form:"value,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

package models

import "time"

//Store ...
type Store struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	Name      string     `json:"name,omitempty" form:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

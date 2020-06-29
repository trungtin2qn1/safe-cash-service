package models

import "time"

//Store ...
type Store struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	Name      string     `json:"name" form:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	District string `json:"district"`
	Ward string `json:"ward"`
	Street string `json:"street"`
}

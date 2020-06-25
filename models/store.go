package models

import "time"

//Store ...
type Store struct {
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	Name      string     `json:"name,omitempty" form:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Country string `json:"country,omitempty"`
	State string `json:"state,omitempty"`
	City string `json:"city,omitempty"`
	District string `json:"district,omitempty"`
	Ward string `json:"ward,omitempty"`
	Street string `json:"street,omitempty"`
}

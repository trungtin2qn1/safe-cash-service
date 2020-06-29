package models

import "time"

//StoreJunctionUser ...
type StoreJunctionUser struct{
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Role string `json:"role"`
	UserID *string `json:"user_id"`
	StoreID *string `json:"store_id"`
}

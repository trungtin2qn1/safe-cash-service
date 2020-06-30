package models

import (
	"time"
)

//RegisterMerchant ...
type RegisterMerchant struct{
	ID        string     `json:"id,omitempty" form:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Content string `json:"content"`
	Email string `json:"email"`
	Name string `json:"name"`
	UserAddress string `json:"user_address"`
	StoreName string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	Status int `json:"status"`
	UserID *string `json:"user_id"`
}
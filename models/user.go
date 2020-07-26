package models

import "time"

// User ...
type User struct {
	ID           string     `json:"id,omitempty"`
	Email        string     `json:"email" form:"email"`
	Password     string     `json:"-" gorm:"password"`
	PhoneNumber  string     `json:"phone_number" form:"phone_number"`
	FirstName    string     `json:"first_name" form:"first_name"`
	LastName     string     `json:"last_name" form:"last_name"`
	Role         string     `json:"role" form:"role"`
	Avatar       string     `json:"avatar" form:"avatar"`
	Token        string     `json:"token,omitempty" form:"token,omitempty" gorm:"-"`
	RefreshToken string     `json:"refresh_token,omitempty" form:"refresh_token,omitempty" gorm:"-"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

package models

import "time"

// User ...
type User struct {
	ID          string     `json:"id,omitempty"`
	Email       string     `json:"email,omitempty" form:"email,omitempty"`
	Password    string     `json:"-" gorm:"password"`
	PhoneNumber string     `json:"phone_number,omitempty" form:"phone_number,omitempty"`
	FirstName   string     `json:"first_name,omitempty" form:"first_name,omitempty"`
	LastName    string     `json:"last_name,omitempty" form:"last_name,omitempty"`
	Position    int        `json:"position,omitempty" form:"position,omitempty"`
	Role        string     `json:"role,omitempty" form:"role,omitempty"`
	Avatar      string     `json:"avatar,omitempty" form:"avatar,omitempty"`
	Token       string     `json:"token,omitempty" form:"token,omitempty" gorm:"-"`
	StoreID     *string    `json:"store_id,omitempty" form:"store_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

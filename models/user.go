package models

// User ...
type User struct {
	ID          string  `json:"id,omitempty"`
	Email       string  `json:"email,omitempty" form:"email,omitempty"`
	Password    string  `json:"-" gorm:"password"`
	PhoneNumber string  `json:"phone_number,omitempty" form:"phone_number,omitempty"`
	FirstName   string  `json:"first_name,omitempty" form:"first_name,omitempty"`
	LastName    string  `json:"last_name,omitempty" form:"last_name,omitempty"`
	Position    int     `json:"position,omitempty" form:"position,omitempty"`
	Token       string  `json:"token,omitempty" form:"token,omitempty" gorm:"-"`
	StoreID     *string `json:"store_id,omitempty" form:"store_id,omitempty"`
}

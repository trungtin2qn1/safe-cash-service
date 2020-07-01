package display

//User...
type User struct {
	ID          string     `json:"id,omitempty"`
	Email       string     `json:"email" form:"email"`
	Password    string     `json:"-" gorm:"password"`
	PhoneNumber string     `json:"phone_number" form:"phone_number"`
	FirstName   string     `json:"first_name" form:"first_name"`
	LastName    string     `json:"last_name" form:"last_name"`
	Position    int        `json:"position" form:"position"`
	Role        string     `json:"role" form:"role"`
	Avatar      string     `json:"avatar" form:"avatar"`
	Token       string     `json:"token,omitempty" form:"token,omitempty" gorm:"-"`
	RefreshToken string `json:"refresh_token,omitempty" form:"refresh_token,omitempty" gorm:"-"`
	StoreID     *string    `json:"store_id" form:"store_id"`
	StoreName      string     `json:"store_name" form:"store_name"`
}

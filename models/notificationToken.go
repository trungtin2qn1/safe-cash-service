package models

//NotificationToken ...
type NotificationToken struct {
	ID     string  `json:"id,omitempty" form:"id,omitempty"`
	Token  string  `json:"token,omitempty" form:"token,omitempty"`
	UserID *string `json:"user_id,omitempty" form:"user_id,omitempty"`
}

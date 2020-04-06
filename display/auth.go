package display

import "safe-cash-service/models"

//AuthRes ...
type AuthRes struct {
	User  models.User  `json:"user,omitempty"`
	Store models.Store `json:"store,omitempty"`
	Token string       `json:"token,omitempty"`
}

package display

import (
	"safe-cash-service/models"
)

//StoreMedia :
type StoreMedia struct {
	Medias    []models.StoreMedia `json:"medias"`
	IsSuccess bool                `json:"is_success"`
	Username  string              `json:"username"`
}

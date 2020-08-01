package display

import (
	"safe-cash-service/models"
	"time"
)

//StoreMedia :
type StoreMedia struct {
	Medias    []models.StoreMedia `json:"medias"`
	IsSuccess bool                `json:"is_success"`
	Username  string              `json:"username"`
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
}

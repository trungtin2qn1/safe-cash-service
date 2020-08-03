package display

import (
	"time"
)

//StoreMedia :
type StoreMedia struct {
	// Medias    []models.StoreMedia `json:"medias"`
	Images    []Image    `json:"images"`
	Videos    []Video    `json:"videos"`
	IsSuccess bool       `json:"is_success"`
	Username  string     `json:"username"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Method    string     `json:"method"`
}

//Image :
type Image struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

//Video :
type Video struct {
	URL string `json:"url"`
}

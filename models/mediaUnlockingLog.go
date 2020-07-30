package models

import "time"

//MediaUnlockingLog ...
type MediaUnlockingLog struct {
	ID             string     `json:"id,omitempty"`
	StoreMediaID   *string    `json:"store_media_id,omitempty" form:"store_media_id,omitempty"`
	UnlockingLogID *string    `json:"unlocking_log_id,omitempty" form:"unlocking_log_id,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}

//TableName :
func (MediaUnlockingLog) TableName() string {
	return "media_unlocking_logs"
}

package models

import "time"

//ClientCredential ...
type ClientCredential struct {
	ID          string     `json:"id,omitempty" form:"id,omitempty"`
	FingerPrint string     `json:"finger_print,omitempty" form:"finger_print,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	StoreID     *string    `json:"store_id,omitempty" form:"store_id,omitempty"`
}

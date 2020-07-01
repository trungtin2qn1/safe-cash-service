package display

import "safe-cash-service/models"

//UnlockingLog ...
type UnlockingLog struct {
	models.UnlockingLog
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

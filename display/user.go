package display

import (
	"safe-cash-service/models"
)

//User :
type User struct {
	models.User
	StoreID   *string `json:"store_id" form:"store_id"`
	StoreName string  `json:"store_name" form:"store_name"`
}

package notification

import (
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//SaveNotificationTokenReq ...
type SaveNotificationTokenReq struct {
	Token string `json:"token,omitempty" form:"token,omitempty"`
}

//SaveToken ...
func SaveToken(userID, token string) (models.NotificationToken, error) {
	notificationToken := models.NotificationToken{
		UserID: &userID,
		Token:  token,
		Value:  token,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&notificationToken)

	return notificationToken, dbConn.Error
}

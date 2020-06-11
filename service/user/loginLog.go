package user

import (
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//CreateLoginLog ...
func CreateLoginLog(userID *string) (*models.UserLoginLog, error) {
	userLoginLog := &models.UserLoginLog{
		UserID: userID,
	}

	dbConn := db.GetDB()
	dbConn = dbConn.Create(&userLoginLog)

	return userLoginLog, dbConn.Error
}

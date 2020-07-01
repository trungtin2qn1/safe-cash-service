package unlockinglog

import (
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//CreateUnlockingLog ...
func CreateUnlockingLog(content, method string, isSuccess bool, userID *string) (models.UnlockingLog, error) {
	unlockingLog := models.UnlockingLog{
		Content:   content,
		IsSuccess: &isSuccess,
		UserID:    userID,
		Method:    method,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&unlockingLog)

	return unlockingLog, dbConn.Error
}

//GetUnlockingLogsByUserID ...
func GetUnlockingLogsByUserID(userID string) ([]models.UnlockingLog, error) {
	res := []models.UnlockingLog{}

	dbConn := db.GetDB()
	dbConn = dbConn.Where("user_id = ?", userID).Find(&res)

	return res, nil
}

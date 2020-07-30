package storemedia

import (
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//CreateMediaUnlockingLog ...
func CreateMediaUnlockingLog(storeMediaID, unlockingLogID *string) (models.MediaUnlockingLog, error) {

	mediaUnlockingLog := models.MediaUnlockingLog{
		StoreMediaID:   storeMediaID,
		UnlockingLogID: unlockingLogID,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&mediaUnlockingLog)

	return mediaUnlockingLog, dbConn.Error
}

//GetMediaUnlockingLogByUnlockingLogID :
func GetMediaUnlockingLogByUnlockingLogID(unlockingLogID string) ([]models.MediaUnlockingLog, error) {
	res := []models.MediaUnlockingLog{}

	dbConn := db.GetDB()

	// if date != "" {
	// 	from := date + " 00:00:00"
	// 	to := date + " 23:59:59"
	// 	dbConn = dbConn.Where("created_at between ? and ?", from, to)
	// }

	dbConn = dbConn.Where("unlocking_log_id = ?", unlockingLogID).Find(&res)

	return res, dbConn.Error
}

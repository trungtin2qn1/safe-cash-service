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

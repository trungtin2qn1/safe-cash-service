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

//GetUnlockingLogByID ...
func GetUnlockingLogByID(id string) (models.UnlockingLog, error) {
	res := models.UnlockingLog{}

	dbConn := db.GetDB()
	dbConn = dbConn.Where("id = ?", id).Find(&res)

	return res, nil
}

//GetUnlockingLogsByUserID ...
func GetUnlockingLogsByUserID(userID string) ([]models.UnlockingLog, error) {
	res := []models.UnlockingLog{}

	dbConn := db.GetDB()
	dbConn = dbConn.
		Order("id desc").
		Where("user_id = ?", userID).Find(&res)

	return res, nil
}

//GetUnlockingLogsByUserIDAndDate ...
func GetUnlockingLogsByUserIDAndDate(userID, date string) ([]models.UnlockingLog, error) {
	res := []models.UnlockingLog{}

	dbConn := db.GetDB()

	if date != "" {
		from := date + " 00:00:00"
		to := date + " 23:59:59"
		dbConn = dbConn.Where("created_at between ? and ?", from, to)
	}

	dbConn = dbConn.
		Order("id desc").
		Where("user_id = ?", userID).Find(&res)

	return res, nil
}

//GetSupport ...
type GetSupport struct {
	Offset int `json:"offset,omitempty" form:"offset,omitempty"`
	Limit  int `json:"limit,omitempty" form:"limit,omitempty"`
}

//getDefault ...
func (support *GetSupport) getDefault() {

	if support.Offset == -1 {
		support.Offset = 0
	}

	if support.Limit <= 0 {
		support.Limit = 15
	}
}

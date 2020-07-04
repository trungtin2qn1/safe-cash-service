package notification

import (
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//UpdateStatus ...
func UpdateStatus(noti *models.Notification, status bool) error {

	dbConn := db.GetDB()
	dbConn = dbConn.Model(&noti).Where("id = ?", noti.ID).Update("is_read", status)

	return dbConn.Error
}

//Create ...
func Create(title, body string, userID *string) (models.Notification, error) {
	notification := models.Notification{
		Title:  title,
		Body:   body,
		UserID: userID,
		IsRead: false,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&notification)

	return notification, nil
}

//GetByUserID ...
func GetByUserID(userID string, support GetSupport) ([]models.Notification, error) {

	notifications := []models.Notification{}

	dbConn := db.GetDB()

	support.getDefault()

	dbConn = dbConn.
		Offset(support.Offset).
		Limit(support.Limit).
		Order("id desc").
		Where("user_id = ?", userID).
		Find(&notifications)

	return notifications, dbConn.Error
}

//GetByID ...
func GetByID(id string) (models.Notification, error) {

	notifications := models.Notification{}

	dbConn := db.GetDB()

	dbConn = dbConn.Where("id = ?", id).Find(&notifications)

	return notifications, dbConn.Error
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
		support.Limit = 10
	}
}

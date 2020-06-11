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
func GetByUserID(userID string) ([]models.Notification, error) {

	notifications := []models.Notification{}

	dbConn := db.GetDB()

	dbConn = dbConn.Where("user_id = ?", userID).Find(&notifications)

	return notifications, dbConn.Error
}

//GetByID ...
func GetByID(id string) (models.Notification, error) {

	notifications := models.Notification{}

	dbConn := db.GetDB()

	dbConn = dbConn.Where("id = ?", id).Find(&notifications)

	return notifications, dbConn.Error
}

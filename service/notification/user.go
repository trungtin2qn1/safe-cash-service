package notification

import (
	"encoding/json"
	"errors"
	"net/http"
	"safe-cash-service/db"
	"safe-cash-service/models"
	"safe-cash-service/utils"
)

//GetOwnerStoreTokenByStoreID :
func GetOwnerStoreTokenByStoreID(storeID string) ([]models.NotificationToken, error) {
	dbConn := db.GetDB()

	notiTokens := []models.NotificationToken{}

	storeJunctionUser := models.StoreJunctionUser{}

	dbConnTemp := dbConn.Where("store_id = ? and role = ?", storeID, "owner").Find(&storeJunctionUser)
	if dbConnTemp.Error != nil {
		return nil, dbConnTemp.Error
	}

	dbConnTemp = dbConn.Where("user_id = ?", storeJunctionUser.UserID).Find(&notiTokens)
	if dbConnTemp.Error != nil {
		return nil, dbConnTemp.Error
	}

	return notiTokens, nil
}

//GetOwnerStoreNotificationByUserID ...
func GetOwnerStoreNotificationByUserID(userID string) ([]models.NotificationToken, error) {
	dbConn := db.GetDB()

	notiTokens := []models.NotificationToken{}

	storeJunctionUser := models.StoreJunctionUser{}

	dbConnTemp := dbConn.Where("user_id = ? and role = ?", userID, "owner").Find(&storeJunctionUser)
	if dbConnTemp.Error != nil {
		return nil, dbConnTemp.Error
	}

	dbConnTemp = dbConn.Where("user_id = ?", storeJunctionUser.UserID).Find(&notiTokens)
	if dbConnTemp.Error != nil {
		return nil, dbConnTemp.Error
	}

	return notiTokens, nil
}

//FCMNotificationRequestData ...
type FCMNotificationRequestData struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

//FCMNotificationRequest ...
type FCMNotificationRequest struct {
	Notification FCMNotificationRequestData `json:"notification,omitempty"`
	To           string                     `json:"to,omitempty"`
}

//Send ...
func Send(notiToken models.NotificationToken) error {

	header := http.Header{}

	header.Set("Content-Type", "application/json")
	header.Set("Authorization", "key=AAAA9z3HQYc:APA91bGO6i6MLR6Pag791H0oa5qFDAMCneaVb4yJZ49SiSAU4zyG5Z5yZzpV7--TDgFvJ9sRHPZ0TvkePhUDdkVzdAHweoUiYdIVhCtI4PocA4-C9U0_CNmRd_aXekL-iLN4V2-1TmrH")

	req := FCMNotificationRequest{
		Notification: FCMNotificationRequestData{
			Title: "Có ai đó vừa mở két tiền của bạn",
			Body:  "Bạn hãy kiểm tra xem ai đó vừa mở két tiền của bạn đấy",
		},
		To: notiToken.Token,
	}

	bodyReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := utils.SendHTTPRequest(header, bodyReq, "https://fcm.googleapis.com/fcm/send", "POST")

	if resp == nil {
		return errors.New("Response is null")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Error in sending message to fcm service")
	}

	return err
}

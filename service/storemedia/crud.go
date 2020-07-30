package storemedia

import (
	"errors"
	"log"
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//CreateStoreMedia ...
func CreateStoreMedia(name, t, thumbnail string, userID, storeID *string) (models.StoreMedia, error) {

	storeMedia := models.StoreMedia{
		UserID:    userID,
		StoreID:   storeID,
		Name:      name,
		Thumbnail: thumbnail,
		Type:      t,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&storeMedia)

	return storeMedia, dbConn.Error
}

//GetByID :
func GetByID(id string) (models.StoreMedia, error) {
	res := models.StoreMedia{}

	dbConn := db.GetDB()

	dbConn = dbConn.Where("id = ?", id).Find(&res)

	return res, dbConn.Error
}

//GetByStoreIDAndDate :
func GetByStoreIDAndDate(storeID, date string) ([]models.StoreMedia, error) {
	res := []models.StoreMedia{}

	dbConn := db.GetDB()

	if date != "" {
		from := date + " 00:00:00"
		to := date + " 23:59:59"
		dbConn = dbConn.Where("created_at between ? and ?", from, to)
	}

	dbConn = dbConn.Where("store_id = ?", storeID).Find(&res)

	return res, dbConn.Error
}

// GetStoreMediaByStoreID ...
func GetStoreMediaByStoreID(storeID string) ([]models.StoreMedia, error) {
	dbConn := db.GetDB()

	storeMedias := []models.StoreMedia{}

	res := dbConn.Where("store_id = ?", storeID).Find(&storeMedias)
	if res.Error != nil {
		log.Println(res.Error)
		return storeMedias, errors.New("Data or data type is invalid")
	}
	return storeMedias, nil
}

// GetStoreMediaByUserID ...
func GetStoreMediaByUserID(userID string) ([]models.StoreMedia, error) {
	dbConn := db.GetDB()

	storeMedias := []models.StoreMedia{}

	res := dbConn.Where("user_id = ?", userID).Find(&storeMedias)
	if res.Error != nil {
		log.Println(res.Error)
		return storeMedias, errors.New("Data or data type is invalid")
	}

	return storeMedias, nil
}

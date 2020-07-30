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

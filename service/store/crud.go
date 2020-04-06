package store

import (
	"errors"
	"log"
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//CreateStore ...
func CreateStore(name string) (models.Store, error) {

	store := models.Store{
		Name: name,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&store)

	return store, dbConn.Error
}

// GetStoreByID ...
func GetStoreByID(id string) (models.Store, error) {
	dbConn := db.GetDB()

	store := models.Store{}

	res := dbConn.Where("id = ?", id).Find(&store)
	if res.Error != nil {
		log.Println(res.Error)
		return store, errors.New("Data or data type is invalid")
	}
	return store, nil
}

// GetStoreByName ...
func GetStoreByName(name string) (models.Store, error) {
	dbConn := db.GetDB()

	store := models.Store{}

	res := dbConn.Where("name = ?", name).First(&store)
	if res.Error != nil {
		log.Println(res.Error)
		return store, errors.New("Data or data type is invalid")
	}

	return store, nil
}

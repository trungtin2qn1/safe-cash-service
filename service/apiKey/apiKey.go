package apiKey

import (
	"cinema-admin/models"
	"errors"
	"safe-cash-service/db"
)

//GetByKey ...
func GetByKey(key string) (models.APIKey, error) {
	dbConn := db.GetDB()

	apiKey := models.APIKey{}
	res := dbConn.Where("value = ?", key).Find(&apiKey)
	if res.Error != nil {
		return apiKey, errors.New("Data or data type is invalid")
	}

	return apiKey, nil
}

package apikey

import (
	"errors"
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//GetByKey ...
func GetByKey(key string) (models.ApiKey, error) {
	dbConn := db.GetDB()

	apiKey := models.ApiKey{}
	res := dbConn.Where("value = ?", key).Find(&apiKey)
	if res.Error != nil {
		return apiKey, errors.New("Data or data type is invalid")
	}

	return apiKey, nil
}

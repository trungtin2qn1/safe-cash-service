package clientcredential

import (
	"safe-cash-service/db"
	"safe-cash-service/models"

	"github.com/google/uuid"
)

//GetByFingerPrintAndStoreID :
func GetByFingerPrintAndStoreID(fingerPrint, storeID string) (*models.ClientCredential, error) {
	clientCredential := &models.ClientCredential{}

	dbConn := db.GetDB()

	dbConn = dbConn.
		Where("store_id = ? and finger_print = ?", storeID, fingerPrint).
		Find(&clientCredential)

	return clientCredential, dbConn.Error
}

//GetByStoreID :
func GetByStoreID(storeID string) (*models.ClientCredential, error) {
	clientCredential := &models.ClientCredential{}

	dbConn := db.GetDB()

	dbConn = dbConn.
		Where("store_id = ? ", storeID).
		Find(&clientCredential)

	return clientCredential, dbConn.Error
}

//CreateByValue ...
func CreateByValue(storeID string) (*models.ClientCredential, error) {

	genID := uuid.New()

	clientCredential := &models.ClientCredential{
		FingerPrint: genID.String(),
		StoreID:     &storeID,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&clientCredential)

	return clientCredential, dbConn.Error

}

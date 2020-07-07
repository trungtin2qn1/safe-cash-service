package clientcredential

import (
	"safe-cash-service/db"
	"safe-cash-service/models"
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

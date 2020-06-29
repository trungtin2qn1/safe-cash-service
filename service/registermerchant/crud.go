package registermerchant

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"safe-cash-service/db"
	"safe-cash-service/models"
)

const PENDINGREQUEST = 0;
const ACCEPTREQUEST = 1;
const REJECTREQUEST = 2;

//RegisterRequest ...
type RegisterRequest struct{
	Content string `json:"content"`
	Email string `json:"email"`
	Name string `json:"name"`
	UserAddress string `json:"user_address"`
	StoreName string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	Status int `json:"status"`
}

//CreateRegisterMerchant ...
func CreateRegisterMerchant(request RegisterRequest, userID *string) (models.RegisterMerchant, error) {

	tempRegisterMerchant, _ := GetRegisterMerchantByUserID(*userID)
	if tempRegisterMerchant.ID != ""{
		return models.RegisterMerchant{}, errors.New("User has registered to be a merchant")
	}

	registerMerchant := models.RegisterMerchant{
		Content: request.Content,
		Email: request.Email,
		Name: request.Name,
		UserAddress: request.UserAddress,
		StoreName: request.StoreName,
		StoreAddress: request.StoreAddress,
		UserID: 	 userID,
		Status: PENDINGREQUEST,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&registerMerchant)

	return registerMerchant, dbConn.Error
}

//GetRegisterMerchantByUserID ...
func GetRegisterMerchantByUserID(userID string) (models.RegisterMerchant, error) {
	dbConn := db.GetDB()

	registerMerchant := models.RegisterMerchant{}
	res := dbConn.Where("user_id = ?", userID).Find(&registerMerchant)
	if res.Error != nil {
		log.Println(res.Error)
		return registerMerchant, errors.New("Data or data type is invalid")
	}

	return registerMerchant, nil
}

//updateInfo ...
func updateInfo(dbConn *gorm.DB, oldInfo models.RegisterMerchant, newInfo *models.RegisterMerchant) error {

	dbConn = dbConn.Model(&oldInfo).Where("id = ?", oldInfo.ID).Update(newInfo)
	return dbConn.Error
}

//UpdateInfo ...
func UpdateInfo(id string, newRegisterMerchant *models.RegisterMerchant) error {

	dbConn := db.GetDB()

	registerMerchant := models.RegisterMerchant{}
	dbConn = dbConn.Where("id = ?", id).Find(&registerMerchant)
	if dbConn.Error != nil {
		log.Println(dbConn.Error)
		return dbConn.Error
	}

	*newRegisterMerchant = registerMerchant
	status := newRegisterMerchant.Status
	newRegisterMerchant.Status = status

	return updateInfo(dbConn, registerMerchant, newRegisterMerchant)

}

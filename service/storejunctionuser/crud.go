package storejunctionuser

import (
	"errors"
	"log"
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//CreateStoreJunctionUser ...
func CreateStoreJunctionUser(role string, userID, storeID *string) (models.StoreJunctionUser, error) {

	storeJunctionUser := models.StoreJunctionUser{
		Role: 		 role,
		UserID: 	 userID,
		StoreID:     storeID,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&storeJunctionUser)

	return storeJunctionUser, dbConn.Error
}


//GetStoreJunctionUserByUserIDAndStoreID ...
func GetStoreJunctionUserByUserIDAndStoreID(userID, storeID string) (models.StoreJunctionUser, error) {
	dbConn := db.GetDB()

	storeJunctionUser := models.StoreJunctionUser{}
	res := dbConn.Where("user_id = ? and store_id = ?", userID, storeID).Find(&storeJunctionUser)
	if res.Error != nil {
		log.Println(res.Error)
		return storeJunctionUser, errors.New("Data or data type is invalid")
	}

	return storeJunctionUser, nil
}

//GetStoreJunctionUsersByStoreID ...
func GetStoreJunctionUsersByStoreID(storeID string) ([]models.StoreJunctionUser, error) {
	dbConn := db.GetDB()

	storeJunctionUsers := []models.StoreJunctionUser{}
	res := dbConn.Where("store_id = ?", storeID).Find(&storeJunctionUsers)
	if res.Error != nil {
		log.Println(res.Error)
		return storeJunctionUsers, errors.New("Data or data type is invalid")
	}

	return storeJunctionUsers, nil
}

//GetStoreJunctionUsersByUserID ...
func GetStoreJunctionUsersByUserID(userID string) ([]models.StoreJunctionUser, error) {
	dbConn := db.GetDB()

	storeJunctionUsers := []models.StoreJunctionUser{}
	res := dbConn.Where("user_id = ?", userID).Find(&storeJunctionUsers)
	if res.Error != nil {
		log.Println(res.Error)
		return storeJunctionUsers, errors.New("Data or data type is invalid")
	}

	return storeJunctionUsers, nil
}

////updateInfo ...
//func updateInfo(dbConn *gorm.DB, oldUserInfo models.User, newUserInfo *models.User) error {
//
//	dbConn = dbConn.Model(&oldUserInfo).Where("id = ?", oldUserInfo.ID).Update(newUserInfo)
//	return dbConn.Error
//}
//
////UpdateInfo ...
//func UpdateInfo(userID string, userInfo *models.User) error {
//
//	dbConn := db.GetDB()
//
//	user := models.User{}
//	dbConn = dbConn.Where("id = ?", userID).Find(&user)
//	if dbConn.Error != nil {
//		log.Println(dbConn.Error)
//		return dbConn.Error
//	}
//
//	userInfo.ID = userID
//	userInfo.Password = user.Password
//	userInfo.Email = user.Email // Optional
//
//	return updateInfo(dbConn, user, userInfo)
//
//}

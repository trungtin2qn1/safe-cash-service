package user

import (
	"errors"
	"log"
	"safe-cash-service/db"
	"safe-cash-service/models"

	"github.com/jinzhu/gorm"
)

//CreateUser ...
func CreateUser(email, password,
	phoneNumber, firstName, lastName string) (models.User, error) {

	user := models.User{
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
		LastName:    lastName,
	}

	dbConn := db.GetDB()

	dbConn = dbConn.Create(&user)

	return user, dbConn.Error
}

// GetUserByID ...
func GetUserByID(id string) (models.User, error) {
	dbConn := db.GetDB()

	user := models.User{}

	res := dbConn.Where("id = ?", id).Find(&user)
	if res.Error != nil {
		log.Println(res.Error)
		return user, errors.New("Data or data type is invalid")
	}
	return user, nil
}

// GetUserByEmail ...
func GetUserByEmail(email string) (models.User, error) {
	dbConn := db.GetDB()

	user := models.User{}
	res := dbConn.Where("email = ?", email).First(&user)
	if res.Error != nil {
		log.Println(res.Error)
		return user, errors.New("Data or data type is invalid")
	}

	return user, nil
}

//GetUsersByStoreID ...
func GetUsersByStoreID(storeID string) ([]models.User, error) {
	dbConn := db.GetDB()

	users := []models.User{}

	storeJunctionUsers := []models.StoreJunctionUser{}

	dbConn = dbConn.Where("store_id = ?", storeID).Find(&storeJunctionUsers)

	for _, v := range storeJunctionUsers {
		user := models.User{}
		dbConn := db.GetDB()
		dbConn = dbConn.Where("id = ?", v.UserID).Find(&user)
		if dbConn.Error != nil {
			return users, dbConn.Error
		}
		user.Role = v.Role
		users = append(users, user)
	}

	return users, dbConn.Error
}

//updateInfo ...
func updateInfo(dbConn *gorm.DB, oldUserInfo models.User, newUserInfo *models.User) error {

	dbConn = dbConn.Model(&oldUserInfo).Where("id = ?", oldUserInfo.ID).Update(newUserInfo)
	return dbConn.Error
}

//UpdateInfo ...
func UpdateInfo(userID string, userInfo *models.User) error {

	dbConn := db.GetDB()

	user := models.User{}
	dbConn = dbConn.Where("id = ?", userID).Find(&user)
	if dbConn.Error != nil {
		log.Println(dbConn.Error)
		return dbConn.Error
	}

	userInfo.ID = userID
	userInfo.Password = user.Password
	userInfo.Email = user.Email // Optional

	return updateInfo(dbConn, user, userInfo)

}

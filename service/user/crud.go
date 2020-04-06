package user

import (
	"errors"
	"log"
	"safe-cash-service/db"
	"safe-cash-service/models"
)

//CreateUser ...
func CreateUser(email, password,
	phoneNumber, firstName, lastName string, position int, storeID *string) (models.User, error) {

	user := models.User{
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
		LastName:    lastName,
		Position:    position,
		StoreID:     storeID,
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

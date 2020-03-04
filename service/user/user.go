package user

import (
	"errors"
	"fmt"
	"log"
	"safe-cash-service/db"
	"safe-cash-service/models"
	"safe-cash-service/utils"
	"safe-cash-service/utils/jwt"
)

//AuthReq ...
type AuthReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func checkAuthData(email string, password string) bool {
	if email == "" {
		return false
	}
	if utils.ValidateFormat(email) != nil {
		return false
	}
	if len(password) < 6 {
		return false
	}
	return true
}

//Login ...
func Login(email string, password string) (models.User, error) {
	user := models.User{}
	var err error
	if !(checkAuthData(email, password)) {
		err = fmt.Errorf("%s", "Email or password is invalid")
		return user, err
	}

	user, err = GetUserByEmail(email)

	if err != nil {
		err = fmt.Errorf("%s", "User is not available")
		return user, err
	}

	check, err := utils.Compare(user.Password, password)

	if err != nil {
		err = fmt.Errorf("%s", "Password is not right")
		return user, err
	}

	if check == false {
		err = fmt.Errorf("%s", "Password is not right")
		return user, err
	}

	token, err := jwt.IssueToken(user.ID, user.Email)
	user.Token = token

	return user, err
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

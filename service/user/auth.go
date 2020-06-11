package user

import (
	"errors"
	"fmt"
	"log"
	"safe-cash-service/db"
	"safe-cash-service/models"
	"safe-cash-service/service/store"
	"safe-cash-service/utils"
	"safe-cash-service/utils/jwt"
)

//AuthReq ...
type AuthReq struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	StoreName string `json:"store_name,omitempty"`
}

func checkAuthData(email string, password string) bool {
	if email == "" {
		log.Println("email is null")
		return false
	}
	if utils.ValidateFormat(email) != nil {
		log.Println("Email's format is not valid")
		return false
	}
	if len(password) < 6 {
		log.Println("Len of password is not greater than 5")
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

	token, err := jwt.IssueToken(user.ID, user.Email, "")
	user.Token = token

	return user, err
}

//Register ...
func Register(email, password string, storeID *string) (models.User, error) {
	//TODO:
	// Get userid and store id from token

	user := models.User{}
	var err error
	if !(checkAuthData(email, password)) {
		err = fmt.Errorf("%s", "Email or password is invalid")
		return user, err
	}

	user, err = GetUserByEmail(email)

	if user.ID != "" {
		err = fmt.Errorf("%s", "User is not available")
		return user, err
	}

	hashPwd, _ := utils.Generate(password)

	user, err = CreateUser(email, hashPwd, "", "", "", 0, storeID)

	if err != nil {
		return user, err
	}

	token, err := jwt.IssueToken(user.ID, user.Email, "")
	user.Token = token

	return user, err
}

//RegisterPublic ...
func RegisterPublic(email, password, storeName string) (models.User, models.Store, error) {
	user := models.User{}
	storeInfo := models.Store{}
	var err error
	if !(checkAuthData(email, password)) {
		err = fmt.Errorf("%s", "Email or password is invalid")
		return user, storeInfo, err
	}

	storeInfo, err = store.GetStoreByName(storeName)

	if storeInfo.ID != "" {
		err = fmt.Errorf("%s", "Store has been registered")
		return user, storeInfo, err
	}

	storeInfo, err = store.CreateStore(storeName)

	if storeInfo.ID == "" {
		err = fmt.Errorf("%s", "Can't add store info to database")
		return user, storeInfo, err
	}

	hashPwd, _ := utils.Generate(password)

	user, err = CreateUser(email, hashPwd, "", "", "", 0, &storeInfo.ID)

	if err != nil {
		return user, storeInfo, err
	}

	token, err := jwt.IssueToken(user.ID, user.Email, storeInfo.ID)

	user.Token = token

	return user, storeInfo, err
}

//UpdatePasswordRequest ...
type UpdatePasswordRequest struct {
	ID          string `json:"id,omitempty" form:"id,omitempty"`
	OldPassword string `json:"old_password,omitempty" form:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty" form:"new_password,omitempty"`
}

//UpdatePassword ...
func UpdatePassword(req UpdatePasswordRequest) error {

	if req.OldPassword == req.NewPassword {
		return errors.New("Old password and new password can not be similar")
	}

	if len(req.NewPassword) < 6 {
		return errors.New("Length of new passworld need to be larged than 6")
	}

	dbConn := db.GetDB()

	user := models.User{}
	res := dbConn.Where("id = ?", req.ID).Find(&user)
	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}

	check, err := utils.Compare(user.Password, req.OldPassword)
	if err != nil {
		log.Println(err)
		return err
	}
	if !check {
		log.Println("Wrong password")
		return errors.New("Wrong password")
	}

	newPassword, err := utils.Generate(req.NewPassword)
	if err != nil {
		return err
	}

	oldUser := user

	user.Password = newPassword
	err = updateInfo(dbConn, oldUser, &user)
	if err != nil {
		return err
	}

	return nil
}

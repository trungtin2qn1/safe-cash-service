package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"safe-cash-service/display"
	"safe-cash-service/models"
	storeService "safe-cash-service/service/store"
	"safe-cash-service/service/storejunctionuser"
	"safe-cash-service/service/user"

	"github.com/gin-gonic/gin"
)

//RegisterPublic ...
func RegisterPublic(c *gin.Context) {
	authReq := user.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	user, _, err := user.RegisterPublic(authReq.Email, authReq.Password, authReq.StoreName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, user)
}

// Register ...
func Register(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	authReq := user.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	user, err := user.RegisterForOwner(authReq.Email, authReq.Password, userID, authReq.StoreID, authReq.Role)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, user)
}

// Login ...
func Login(c *gin.Context) {
	authReq := user.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	userInfo, storeInfo, err := user.Login(authReq.Email, authReq.Password, authReq.StoreName)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	_, err = user.CreateLoginLog(&userInfo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	userDisplay := display.User{
		StoreID:   &storeInfo.ID,
		StoreName: storeInfo.Name,
	}

	data, err := json.Marshal(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	err = json.Unmarshal(data, &userDisplay)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, userDisplay)
}

//GetInfo ...
func GetInfo(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)
	interfaceStoreID, _ := c.Get("store_id")
	storeID := fmt.Sprintf("%v", interfaceStoreID)
	user, err := user.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	store, err := storeService.GetStoreByID(storeID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	storeJunctionUser, err := storejunctionuser.GetStoreJunctionUserByUserIDAndStoreID(userID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	userDisplay := display.User{
		StoreID:   &storeID,
		StoreName: store.Name,
	}

	data, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	err = json.Unmarshal(data, &userDisplay)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	userDisplay.Role = storeJunctionUser.Role
	userDisplay.Avatar = Host + userDisplay.Avatar

	c.JSON(200, userDisplay)
}

//UpdateInfo ...
func UpdateInfo(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	newUserInfo := models.User{}
	err := c.ShouldBind(&newUserInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	err = user.UpdateInfo(userID, &newUserInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, newUserInfo)
}

//UpdateInfoV1 ...
func UpdateInfoV1(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	newUserInfo := models.User{}
	err := c.ShouldBind(&newUserInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	avatar, err := handleFormFile(c, "avatar")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	if avatar != "" {
		newUserInfo.Avatar = avatar
	}

	err = user.UpdateInfo(userID, &newUserInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	newUserInfo.Avatar = Host + newUserInfo.Avatar
	c.JSON(200, newUserInfo)
}

//UpdatePassword ...
func UpdatePassword(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	req := user.UpdatePasswordRequest{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	req.ID = userID
	err = user.UpdatePassword(req)
	log.Println("err outside:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}

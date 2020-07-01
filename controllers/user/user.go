package user

import (
	"fmt"
	"log"
	"net/http"
	"safe-cash-service/display"
	"safe-cash-service/models"
	storeService "safe-cash-service/service/store"
	"safe-cash-service/service/user"

	"github.com/gin-gonic/gin"
)

//RegisterPublicV1 ...
func RegisterPublicV1(c *gin.Context) {
	authReq := user.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	user, err := user.RegisterPublicV1(authReq.Email, authReq.Password, authReq.FirstName, authReq.LastName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, user)
}

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
	//storeID := fmt.Sprintf("%v", interfaceStoreID)
	userID := fmt.Sprintf("$v", interfaceUserID)

	authReq := user.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	user, err := user.RegisterForOwner(authReq.Email, authReq.Password, userID, authReq.StoreID)

	if err != nil {
		c.JSON(http.StatusNotAcceptable , gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, user)
}

// LoginV1 ...
func LoginV1(c *gin.Context) {
	authReq := user.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	userInfo, err := user.LoginV1(authReq.Email, authReq.Password)
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

	c.JSON(200, userInfo)
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

	userInfo, err := user.Login(authReq.Email, authReq.Password, authReq.StoreName)
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

	c.JSON(200, userInfo)
}

//GetInfo ...
func GetInfo(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	user, err := user.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}


	storeID := ""
	storeName := ""

	if user.StoreID != nil {
		storeID = *user.StoreID

		store, err := storeService.GetStoreByID(storeID)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": fmt.Sprintf("%s", err),
			})
			return
		}
		storeID = store.ID
		storeName = store.Name
	}

	userDisplay := display.User{
		ID: user.ID,
		StoreID: &storeID,
		StoreName: storeName,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Position: user.Position,
		Role: user.Role,
		Avatar: user.Avatar,
		Token: user.Token,
		RefreshToken: user.RefreshToken,
		PhoneNumber: user.PhoneNumber,
	}

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

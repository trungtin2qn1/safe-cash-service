package user

import (
	"fmt"
	"log"
	"net/http"
	"safe-cash-service/models"
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
	interfaceStoreID, _ := c.Get("store_id")
	storeID := fmt.Sprintf("%v", interfaceStoreID)

	authReq := user.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	user, err := user.Register(authReq.Email, authReq.Password, &storeID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

	user, err := user.Login(authReq.Email, authReq.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, user)
}

//GetInfo ...
func GetInfo(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	user, err := user.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, user)
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

	err = user.UpdateInfo(userID, newUserInfo)
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

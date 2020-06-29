package store

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"safe-cash-service/service/registermerchant"
)

//RegisterMerchant ...
func RegisterMerchant(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	request := registermerchant.RegisterRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	registerMerchant, err := registermerchant.CreateRegisterMerchant(request, &userID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, registerMerchant)
}

//GetRegisterMerchantRequest ...
func GetRegisterMerchantRequest(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	registerMerchant, err := registermerchant.GetRegisterMerchantByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, registerMerchant)
}
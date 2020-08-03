package user

import (
	"net/http"
	"safe-cash-service/service/notification"

	"github.com/gin-gonic/gin"
)

//LogOutRequest :
type LogOutRequest struct {
	FCMToken string `json:"fcm_token" form:"fcm_token"`
}

//LogOut :
func LogOut(c *gin.Context) {

	request := LogOutRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	err = notification.DeleteByValue(request.FCMToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't delete fcm token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}

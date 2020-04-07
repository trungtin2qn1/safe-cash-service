package notification

import (
	"fmt"
	"net/http"
	"safe-cash-service/service/notification"

	"github.com/gin-gonic/gin"
)

//SaveToken ...
func SaveToken(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	req := notification.SaveNotificationTokenReq{}

	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	notificationToken, err := notification.SaveToken(userID, req.Token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, notificationToken)

}

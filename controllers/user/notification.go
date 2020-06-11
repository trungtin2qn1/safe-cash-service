package user

import (
	"fmt"
	"net/http"
	"safe-cash-service/service/notification"

	"github.com/gin-gonic/gin"
)

//GetNotifications ...
func GetNotifications(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	notifications, err := notification.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
	}

	c.JSON(200, notifications)

}

//UpdateNotificationStatus ...
func UpdateNotificationStatus(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	notificationID := c.Param("notification_id")

	noti, err := notification.GetByID(notificationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Notification is not available",
		})
		return
	}

	if *noti.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "This notification is not belong to user",
		})
		return
	}

	err = notification.UpdateStatus(&noti, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server is busy",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}

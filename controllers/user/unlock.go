package user

import (
	"fmt"
	"log"
	"net/http"
	"safe-cash-service/models"
	"safe-cash-service/service/notification"
	"safe-cash-service/service/unlockinglog"
	"safe-cash-service/service/user"

	"github.com/gin-gonic/gin"
)

//Unlock ...
func Unlock(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	unlockingLog := models.UnlockingLog{}

	err := c.ShouldBind(&unlockingLog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	userInfo, err := user.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	//Create unlocking log
	unlockingLogInfo, err := unlockinglog.CreateUnlockingLog(unlockingLog.Content, *unlockingLog.IsSuccess, &userInfo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	// Send notification to mobile app
	go func(userID string) {
		notiTokens, err := notification.GetOwnerStoreNotificationByUserID(userID)
		if err != nil {
			log.Println(err)
			return
		}

		for _, notiToken := range notiTokens {
			err := notification.Send(notiToken)
			if err != nil {
				continue
			}
		}

		_, err = notification.Create("", "", &userID)
		if err != nil {
			log.Println(err)
			return
		}
	}(userInfo.ID)

	c.JSON(200, unlockingLogInfo)

}

//ListUnlockingLogs ...
func ListUnlockingLogs(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	storeID := c.Query("store_id")

	unlockingLogs := []models.UnlockingLog{}

	userInfo, err := user.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Can't find user",
		})
		return
	}

	if storeID != "" {

		if *userInfo.StoreID != storeID {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "User don't belong to this store",
			})
			return
		}

		users, err := user.GetUsersByStoreID(storeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server is busy",
			})
			return
		}

		for _, v := range users {
			logs, err := unlockinglog.GetUnlockingLogsByUserID(v.ID)
			if err != nil {
				continue
			}
			unlockingLogs = append(unlockingLogs, logs...)
		}

	} else {
		unlockingLogs, err = unlockinglog.GetUnlockingLogsByUserID(userID)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "Server is busy",
			})
			return
		}
	}

	c.JSON(200, unlockingLogs)
}

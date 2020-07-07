package middleware

import (
	"net/http"
	"safe-cash-service/service/clientcredential"
	"safe-cash-service/service/storejunctionuser"

	"github.com/gin-gonic/gin"
)

//VerifyClientCridentials ...
func VerifyClientCridentials(c *gin.Context) {

	fingerPrint := c.Request.Header.Get("finger_print")

	userID := c.Request.Header.Get("user_id")

	storeJunctionUser, err := storejunctionuser.GetStoreJunctionUsersByUserID(userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong Finger Print",
		})
		return
	}

	if len(storeJunctionUser) != 1 && storeJunctionUser[0].StoreID != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Database was not synced",
		})
		return
	}

	clientCridential, err := clientcredential.GetByFingerPrintAndStoreID(fingerPrint, *storeJunctionUser[0].StoreID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong Finger Print",
		})
		return
	}
	if clientCridential.ID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong Finger Print",
		})
		return
	}

	c.Next()
}

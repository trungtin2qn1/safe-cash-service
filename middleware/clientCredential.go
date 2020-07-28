package middleware

import (
	"net/http"
	"safe-cash-service/service/clientcredential"

	"github.com/gin-gonic/gin"
)

//VerifyClientCridentials ...
func VerifyClientCridentials(c *gin.Context) {

	storeCredential := c.Request.Header.Get("store_credential")
	fingerPrint := c.Request.Header.Get("finger_print")

	if fingerPrint == "" {
		fingerPrint = storeCredential
	}

	// userID := c.Request.Header.Get("user_id")

	// storeJunctionUser, err := storejunctionuser.GetStoreJunctionUsersByUserID(userID)

	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"message": "Wrong Finger Print",
	// 	})
	// 	return
	// }

	// if len(storeJunctionUser) != 1 && storeJunctionUser[0].StoreID != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Database was not synced",
	// 	})
	// 	return
	// }

	clientCridential, err := clientcredential.GetByFingerPrint(fingerPrint)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong Finger Print",
		})
		return
	}
	if clientCridential.ID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong Finger Print",
		})
		return
	}

	c.Next()
}

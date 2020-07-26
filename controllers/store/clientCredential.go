package store

import (
	"fmt"
	"net/http"
	"safe-cash-service/service/clientcredential"

	"github.com/gin-gonic/gin"
)

//GetFingerPrint :
func GetFingerPrint(c *gin.Context) {
	interfaceStoreID, _ := c.Get("store_id")
	storeID := fmt.Sprintf("%v", interfaceStoreID)

	clientCredential, err := clientcredential.GetByStoreID(storeID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Can't find client credential",
		})
		return
	}

	clientCredential.StoreCredential = clientCredential.FingerPrint
	c.JSON(http.StatusOK, clientCredential)
}

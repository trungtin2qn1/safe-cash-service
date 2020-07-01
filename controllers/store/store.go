package store

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"safe-cash-service/models"
	storeService "safe-cash-service/service/store"
	"safe-cash-service/service/storejunctionuser"
)

//GetALlStoresByUserID ...
func GetAllStoresByUserID(c *gin.Context){
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	storeJunctionUsers, err := storejunctionuser.GetStoreJunctionUsersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	stores := []models.Store{}
	for _, v := range storeJunctionUsers{
		store, err := storeService.GetStoreByID(*v.StoreID)
		if err != nil {
			continue
		}
		stores = append(stores, store)
	}

	c.JSON(200, stores)
}

//GetByID ...
func GetByID(c *gin.Context){
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)
	storeID := c.Param("store_id")

	storeJunctionUser, err := storejunctionuser.GetStoreJunctionUserByUserIDAndStoreID(userID, storeID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	if storeJunctionUser.ID == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "You don't have permission to access this resource",
		})
		return
	}

	store, err := storeService.GetStoreByID(storeID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, store)
}
package merchant

import (
	"fmt"
	"net/http"
	userService "safe-cash-service/service/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//GetStaffsInStore ...
func GetStaffsInStore(c *gin.Context) {
	// interfaceUserID, _ := c.Get("user_id")
	// userID := fmt.Sprintf("%v", interfaceUserID)
	interfaceStoreID, _ := c.Get("store_id")
	storeID := fmt.Sprintf("%v", interfaceStoreID)

	staffs, err := userService.GetUsersByStoreID(storeID)

	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server is busy",
		})
		return
	}

	c.JSON(200, staffs)
}

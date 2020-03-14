package media

import (
	"net/http"
	"safe-cash-service/file"

	"github.com/gin-gonic/gin"
)

//Upload ...
func Upload(c *gin.Context) {
	err := file.Upload(c.Writer, c.Request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

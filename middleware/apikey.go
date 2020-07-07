package middleware

import (
	"log"
	"net/http"
	"safe-cash-service/service/apikey"

	"github.com/gin-gonic/gin"
)

//VerifyAPIKey ...
func VerifyAPIKey(c *gin.Context) {

	key := c.Request.Header.Get("x-api-key")
	log.Println("key:", key)
	apiKey, err := apikey.GetByKey(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong API key",
		})
		return
	}
	if apiKey.ID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong API key",
		})
		return
	}

	c.Next()
}

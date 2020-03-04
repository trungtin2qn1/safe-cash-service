package middleware

import (
	"fmt"
	"net/http"
	"safe-cash-service/service/apiKey"
	"safe-cash-service/utils/jwt"

	"github.com/gin-gonic/gin"
)

//VerifyJWTToken ...
func VerifyJWTToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	rawToken := string(token[len("Tin "):])
	userID, _, err := jwt.VerificationToken(rawToken)
	if err != nil {
		//go utils.LogErrToFile(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	fmt.Println(userID)
	c.Set("user_id", userID)
	c.Next()
}

//CheckAPIKey ...
func CheckAPIKey(c *gin.Context) {
	key := c.Request.Header.Get("api-key")

	if key == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	serviceKey, err := apiKey.GetByKey(key)
	if err != nil {
		//go utils.LogErrToFile(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c.Set("service-key-id", serviceKey.ID)
	c.Next()
}

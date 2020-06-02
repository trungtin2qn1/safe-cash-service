package middleware

import "github.com/gin-gonic/gin"

func VerifyAPIKey(c *gin.Context) {

	c.Next()
}

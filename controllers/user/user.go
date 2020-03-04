package user

import (
	"net/http"
	"safe-cash-service/service/user"

	"github.com/gin-gonic/gin"
)

// Register ...
func Register(c *gin.Context) {

}

// Login ...
func Login(c *gin.Context) {
	authReq := user.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	user, err := user.Login(authReq.Email, authReq.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server is busy",
		})
		return
	}

	c.JSON(200, user)
}

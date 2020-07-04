package user

import (
	"fmt"
	"net/http"
	"safe-cash-service/service/user"
	"safe-cash-service/utils/jwt"
	"time"

	"github.com/gin-gonic/gin"
)

//GetAccessToken ...
func GetAccessToken(c *gin.Context) {
	req := user.AccessTokenRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	rawToken := string(req.RefreshToken[len("Tin "):])
	userID, email, storeID, err := jwt.VerificationToken(rawToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	userInfo, err := user.GetUserByID(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	refreshToken, err := jwt.IssueToken(userID, email, storeID, time.Second*604800)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	accessToken, err := jwt.IssueToken(userID, email, storeID, time.Second*86400)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	userInfo.RefreshToken = refreshToken
	userInfo.Token = accessToken

	c.JSON(http.StatusOK, userInfo)
}

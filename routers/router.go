package routers

import (
	"safe-cash-service/controllers/user"
	"safe-cash-service/middleware"

	"github.com/gin-gonic/gin"
)

// User handler
// Video handler
// Notification handler
//

//SetUpRouter ...
func SetUpRouter() {
	router := gin.Default()

	api := router.Group("/apis")
	api.POST("/login", user.Login)

	auth := api.Group("/auth")
	{
		auth.Use(middleware.VerifyJWTToken)
		auth.POST("/register", user.Register)

		auth.GET("/")
	}

	router.Run(":6000")
}

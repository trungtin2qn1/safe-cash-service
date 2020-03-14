package routers

import (
	mediaController "safe-cash-service/controllers/media"
	"safe-cash-service/controllers/notification"
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
		auth.GET("/self", user.GetInfo)
		auth.PUT("/self", user.UpdateInfo)

		media := auth.Group("/media")
		{
			media.POST("", mediaController.Upload)
		}

		auth.POST("/notification/token", notification.SaveToken)
	}

	router.Run(":6000")
}

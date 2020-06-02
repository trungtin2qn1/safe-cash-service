package routers

import (
	mediaController "safe-cash-service/controllers/media"
	merchantController "safe-cash-service/controllers/merchant"
	notificationController "safe-cash-service/controllers/notification"
	userController "safe-cash-service/controllers/user"
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

	services := router.Group("/services")
	{
		services.Use(middleware.VerifyAPIKey)
		admin := services.Group("/admin")
		{
			admin.POST("/notification", notificationController.SendByAdmin)
		}

		faceRecognition := services.Group("/face-recognition")
		{
			faceRecognition.GET("/user", userController.GeFromExternalService)
		}
	}

	api := router.Group("/apis")
	api.POST("/register", userController.RegisterPublic)
	api.POST("/login", userController.Login)

	auth := api.Group("/auth")
	{
		auth.Use(middleware.VerifyJWTToken)
		auth.POST("/register", userController.Register)
		auth.GET("/self", userController.GetInfo)
		auth.PUT("/self", userController.UpdateInfo)

		media := auth.Group("/media")
		{
			media.POST("", mediaController.Upload)
		}

		auth.POST("/notification/token", notificationController.SaveToken)
		auth.POST("/notification", notificationController.Send)

		merchant := auth.Group("/merchant")
		{
			merchant.GET("/staffs", merchantController.GetStaffsInStore)
		}
	}

	router.Run(":6000")
}

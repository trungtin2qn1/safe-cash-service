package routers

import (
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
			faceRecognition.GET("/user", userController.GetFromExternalService)
		}
	}

	api := router.Group("/apis")
	api.POST("/register", userController.RegisterPublic)
	api.POST("/login", userController.Login)

	auth := api.Group("/auth")
	{
		auth.Use(middleware.VerifyJWTToken)
		auth.POST("/register", userController.Register)
		self := auth.Group("/self")
		{
			self.GET("", userController.GetInfo)
			self.PUT("", userController.UpdateInfo)
			self.PUT("/password", userController.UpdatePassword)
		}

		auth.POST("/unlock", userController.Unlock)
		auth.GET("/notifications", userController.GetNotifications)
		auth.GET("/unlock/logs", userController.ListUnlockingLogs)
		auth.PUT("/notification/:notification_id", userController.UpdateNotificationStatus)
		auth.GET("/staffs", merchantController.GetStaffsInStore)
		auth.POST("/train-model") // Not implemented yet

		auth.POST("/notification/token", notificationController.SaveToken)
		auth.POST("/notification", notificationController.Send)
	}

	router.Run(":6000")
}

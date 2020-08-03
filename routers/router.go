package routers

import (
	"net/http"
	merchantController "safe-cash-service/controllers/merchant"
	notificationController "safe-cash-service/controllers/notification"
	storeController "safe-cash-service/controllers/store"
	storeMediaController "safe-cash-service/controllers/storemedia"
	userController "safe-cash-service/controllers/user"
	"safe-cash-service/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// User handler
// Video handler
// Notification handler
//

//SetUpRouter ...
func SetUpRouter() {
	router := gin.Default()

	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true

	routerConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(routerConfig))

	services := router.Group("/services")
	{
		// services.Use(middleware.VerifyAPIKey)
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
	{

		api.StaticFS("/public", http.Dir("static"))

		services := api.Group("/services")
		{
			services.Use(middleware.VerifyClientCridentials)
			services.POST("/unlock", userController.UnlockByService)
			services.POST("/store/:store_id/medias", storeMediaController.Upload)
		}

		api.POST("/register", userController.RegisterPublic)
		api.POST("/login", userController.Login)
		api.POST("/access-token", userController.GetAccessToken)

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

			auth.GET("/finger-print", storeController.GetFingerPrint)
			auth.GET("/store/:store_id", storeController.GetByID)
			auth.GET("/store/:store_id/medias", storeMediaController.GetByStoreID)
			auth.GET("/store/:store_id/unlock/:unlock_id/medias", storeMediaController.GetByUnlockingID)
			auth.GET("/stores", storeController.GetAllStoresByUserID)
			auth.POST("/unlock", userController.Unlock)
			auth.GET("/notifications", userController.GetNotifications)
			auth.GET("/unlock/logs", userController.ListUnlockingLogs)
			auth.PUT("/notification/:notification_id", userController.UpdateNotificationStatus)
			auth.GET("/staffs", merchantController.GetStaffsInStore)
			auth.POST("/train-model") // Not implemented yet
			auth.POST("/log-out", userController.LogOut)

			auth.POST("/notification/token", notificationController.SaveToken)
			auth.POST("/notification", notificationController.Send)
		}
	}

	router.StaticFS("/public", http.Dir("static"))

	err := router.Run(":5000")
	if err != nil {
		panic(err)
	}
}

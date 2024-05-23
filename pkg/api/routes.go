package api

import (
	"dating-app/internal/handlers"
	"dating-app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/api/v1/signup", handlers.SignUp)
	router.POST("/api/v1/login", handlers.Login)

	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/stack-profiles", handlers.GetUnswipedProfiles)
		protected.GET("/liked-profiles", handlers.ViewLikedProfiles)
		protected.POST("/profiles", handlers.CreateProfile)
		protected.PUT("/profiles", handlers.UpdateProfile)
		protected.POST("/swipe", handlers.SwipeProfile)
		protected.POST("/premium", handlers.PurchasePremium)
	}
}

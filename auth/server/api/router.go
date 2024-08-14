package api

import (
	"net/http"

	"github.com/TheoremN1/Coins/auth/server/api/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	v1 := router.Group("/api/v1")
	{
		userController := controllers.GetUserController()
		users := v1.Group("/users")
		{
			users.GET("/:id", userController.GetUserById)
			users.GET("", userController.GetAllUsers)
			users.DELETE("/:id", userController.DeleteUserById)
		}

		authController := controllers.GetAuthController()
		v1.POST("/registration", authController.Registration)
		v1.POST("/authorization", authController.Authorization)
		v1.PUT("/refresh", authController.RefreshToken)
		v1.GET("/data", authController.ExtractData)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route is not found"})
	})
}

func GetRouter() *gin.Engine {
	return router
}

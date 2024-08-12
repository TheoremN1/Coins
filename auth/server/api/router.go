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
		UserController := controllers.GetUserController()
		user := v1.Group("/user")
		{
			user.GET(":id", UserController.GetId)
			user.GET("", UserController.GetAll)
			user.POST("", UserController.Post)
			user.PUT("", UserController.Put)
			user.DELETE("", UserController.Delete)
		}
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route is not found"})
	})
}

func GetRouter() *gin.Engine {
	return router
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "pong",
			},
		)
	})
	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

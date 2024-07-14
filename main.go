package main

import (
	"github.com/TheoremN1/Coins/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := server.NewRouter(gin.Default())
	router.Run()
}

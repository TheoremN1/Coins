package main

import (
	"github.com/TheoremN1/Coins/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := internal.NewRouter(gin.Default())
	router.Run()
}

package main

import (
	"github.com/TheoremN1/Coins/server"
)

func main() {
	router := server.BuildRouter()
	router.Run()
}

package main

import (
	"github.com/TheoremN1/Coins/server"
)

func main() {
	host := server.NewHost()
	router := host.Router
	router.Run()
}

package main

import RequestsService "github.com/TheoremN1/Coins/RequestsService"

func main() {
	router := RequestsService.NewRouter()
	router.Run()
}

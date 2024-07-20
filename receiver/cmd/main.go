package main

import "github.com/TheoremN1/Coins/receiver"

func main() {
	router := receiver.BuildRouter()
	router.Run()
}

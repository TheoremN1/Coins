package main

import "github.com/TheoremN1/Coins/auth/server"

func main() {
	server := server.GetServer()
	server.Run()
}

package main

import users_service "github.com/TheoremN1/Coins/UsersService"

func main() {
	router := users_service.NewRouter()
	router.Run()
}

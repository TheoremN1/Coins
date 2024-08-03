package main

import (
	"log"
	"os"

	users_service "github.com/TheoremN1/Coins/UsersService"
	"github.com/joho/godotenv"
)

func init() {
	if len(os.Args) > 1 && os.Args[1] == "--nodocker" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	router := users_service.NewRouter()
	router.Run()
}

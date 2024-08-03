package main

import (
	"log"
	"os"

	RequestsService "github.com/TheoremN1/Coins/RequestsService"
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
	router := RequestsService.NewRouter()
	router.Run()
}

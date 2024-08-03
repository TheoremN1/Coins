package main

import (
	"log"
	"os"

	ProductService "github.com/TheoremN1/Coins/ProductsService"
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
	router := ProductService.NewRouter()
	router.Run()
}

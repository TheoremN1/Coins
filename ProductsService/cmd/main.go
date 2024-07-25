package main

import ProductService "github.com/TheoremN1/Coins/ProductsService"

func main() {
	router := ProductService.NewRouter()
	router.Run()
}

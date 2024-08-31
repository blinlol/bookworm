package main

import (
	"log"
	"os"

	"github.com/blinlol/bookworm/cmd/api/web"
)

func main() {
	address := os.Getenv("API_PORT")
	if address == "" {
		address = "8081"
		log.Println("API_PORT doesn't set. Use port =", address)
	}
	address = ":" + address
	router := web.CreateRouter()

	// add middlewares
	router.Use(web.CORSMiddleware())

	// add routes
	router = web.BookRoutes(router)

	// run server
	err := router.Run(address)
	if err != nil {
		log.Fatalln(err)
	}
}

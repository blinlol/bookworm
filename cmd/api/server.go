package main

import (
	"os"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/blinlol/bookworm/cmd/api/web"
    // "bookworm/api/web"
)

func main(){
	address := os.Getenv("API_PORT")
	if address == "" {
		address = "8081"
		log.Println("API_PORT doesn't set. Use port =", address)
	}
	address = ":" + address
	router := gin.Default()

	// add middlewares
	router.Use(web.CORSMiddleware())

	// add routes
	router.GET("/ping", web.Pong)
	router = web.BookRoutes(router)

	// run server
	err := router.Run(address)
	if err != nil {
		log.Fatalln(err)
	}
}
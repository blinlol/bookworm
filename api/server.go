package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/blinlol/bookworm/api/web"
)

func main(){
	address := "localhost:8081"
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
package web

import (
	_"net/http"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Pong)
	return r
}


func BookRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/books", GetBooks)
	r.POST("/books/add", AddBook)
	r.DELETE("/book/:id", DeleteBook)
	return r
}


func StartServer(address string) error {
	router := SetupRouter()
	router = BookRoutes(router)
	return router.Run(address)
}
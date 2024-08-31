package web

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", Pong)
	return router
}

func BookRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/api/books", GetBooks)
	r.POST("/api/books/add", AddBook)

	r.GET("/api/book/:id", GetBook)
	r.DELETE("/api/book/:id", DeleteBook)
	return r
}

package web

import (
	_"net/http"

	"github.com/gin-gonic/gin"
)



func BookRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/api/books", GetBooks)
	r.POST("/api/books/add", AddBook)
	r.DELETE("/api/book/:id", DeleteBook)
	return r
}

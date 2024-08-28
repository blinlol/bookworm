package web

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/blinlol/bookworm/data/dao"
)


var logger *zap.Logger


func GetBooks(c *gin.Context){
	books := dao.AllBooks()
	if books == nil {
		books = make([]*dao.Book, 0)
	}
	c.IndentedJSON(
		http.StatusOK,
		books,
	)
}


func AddBook(c *gin.Context){
	var book dao.Book
	err := json.NewDecoder(c.Request.Body).Decode(&book)
	if err != nil {
		logger.Sugar().Errorln(err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": err},
		)
	}

	b := dao.AddBook(book.Title, book.Author)
	if b != nil {
		c.JSON(
			http.StatusOK,
			*b,
		)
	}
}


func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	dao.DeleteBookById(id)
	c.JSON(http.StatusOK, gin.H{})
}


func Pong(c *gin.Context){
	c.JSON(
		http.StatusOK,
		gin.H{"message": "pong"},
	)
}


func initLogger(){
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
}


func init(){
	initLogger()
}
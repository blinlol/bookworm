package web

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/blinlol/bookworm/model"
	"github.com/blinlol/bookworm/model/dao"
	// "bookworm/model"
    // "bookworm/model/dao"
)


var logger *zap.Logger


func GetBooks(c *gin.Context){
	books := dao.AllBooks()
	if books == nil {
		books = make([]*model.Book, 0)
	}
	body := model.GetBooksResponse{Books: books}
	c.IndentedJSON(
		http.StatusOK,
		body,
	)
}


func AddBook(c *gin.Context){
	var req model.AddBookRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		logger.Sugar().Errorln(err)
		c.JSON(
			http.StatusBadRequest,
			model.ErrorResponse{Message: err.Error()},
		)
	}

	b := dao.AddBook(req.Book.Title, req.Book.Author)
	if b != nil {
		c.JSON(
			http.StatusOK,
			model.AddBookResponse{Book: *b},
		)
	}
}


func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	dao.DeleteBookById(id)
	c.JSON(http.StatusOK, gin.H{})
}


func Pong(c *gin.Context){
	pong := model.PingResponse{Message: "pong"}
	c.JSON(
		http.StatusOK,
		pong,
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
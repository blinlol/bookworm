package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/blinlol/bookworm/model"
	"github.com/blinlol/bookworm/model/dao"
)

var logger *zap.Logger

func GetBooks(c *gin.Context) {
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

func AddBook(c *gin.Context) {
	var req model.AddBookRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		logger.Sugar().Errorln(err)
		c.JSON(
			http.StatusBadRequest,
			model.ErrorResponse{Message: err.Error()},
		)
	}

	b := dao.AddBook(&req.Book)
	if b != nil {
		c.JSON(
			http.StatusOK,
			model.AddBookResponse{Book: *b},
		)
	}
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book := dao.FindBookById(id)
	if book == nil {
		c.JSON(
			http.StatusBadRequest,
			model.ErrorResponse{
				Message: "Book with such id not found",
			},
		)
	} else {
		resp := model.GetBookResponse{
			Book: book,
		}
		c.JSON(
			http.StatusOK,
			resp,
		)
	}
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	dao.DeleteBookById(id)
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var req model.UpdateBookRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		logger.Sugar().Infoln(err)
		c.JSON(
			http.StatusBadRequest,
			model.ErrorResponse{Message: fmt.Sprintf("%v", err)},
		)
	} else if id != req.Book.Id {
		message := "id param and req.Book.Id not equal"
		logger.Sugar().Infoln(message)
		c.JSON(
			http.StatusBadRequest,
			model.ErrorResponse{Message: message},
		)
	} else {
		upd_b := dao.UpdateBook(&req.Book)
		if upd_b == nil {
			message := "error while book updating"
			logger.Sugar().Infoln(message)
			c.JSON(
				http.StatusBadRequest,
				model.ErrorResponse{Message: message},
			)
		} else {
			c.Status(http.StatusOK)
		}
	}
}

func Pong(c *gin.Context) {
	pong := model.PingResponse{Message: "pong"}
	c.JSON(
		http.StatusOK,
		pong,
	)
}

func initLogger() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	initLogger()
}

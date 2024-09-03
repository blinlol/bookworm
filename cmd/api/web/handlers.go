package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/blinlol/bookworm/model"
	"github.com/blinlol/bookworm/model/dao"
	"github.com/blinlol/bookworm/utils"
)

var logger *zap.Logger

func GetBooks(c *gin.Context) {
	ctx := handlerContext()
	books := dao.AllBooks(ctx)
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

	ctx := handlerContext()
	b := dao.AddBook(ctx, req.Book)
	if b != nil {
		c.JSON(
			http.StatusCreated,
			model.AddBookResponse{Book: *b},
		)
	}
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	ctx := handlerContext()
	book := dao.FindBookById(ctx, id)
	if book == nil {
		c.JSON(
			http.StatusNoContent,
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
	ctx := handlerContext()
	dao.DeleteBookById(ctx, id)
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
		ctx := handlerContext()
		success := dao.UpdateBook(ctx, req.Book)
		if success {
			c.Status(http.StatusOK)
		} else {
			message := "error while book updating"
			logger.Sugar().Infoln(message)
			c.JSON(
				http.StatusBadRequest,
				model.ErrorResponse{Message: message},
			)
		}
	}
}

func ParseQuotes(c *gin.Context) {
	var req model.ParseQuotesRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		logger.Sugar().Infoln(err)
		c.JSON(
			http.StatusBadRequest,
			model.ErrorResponse{Message: fmt.Sprintf("%v", err)},
		)
	} else {
		ctx := handlerContext()
		quotes := utils.ParseQuotes(req.Text, req.Separator)
		success := dao.AddQuotes(ctx, req.BookId, quotes)
		if success {
			c.Status(http.StatusOK)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				model.ErrorResponse{Message: fmt.Sprintf("%v", err)},
			)
		}
	}
}

func GetQuotes(c *gin.Context) {
	var res model.GetQuotesResponse
	bookId := c.Param("book_id")
	ctx := handlerContext()
	quotes := dao.GetQuotesByBookId(ctx, bookId)
	if quotes == nil {
		c.JSON(
			http.StatusNoContent,
			model.ErrorResponse{Message: "There no quotes"},
		)
	} else {
		res.BookId = bookId
		res.Quotes = quotes
		c.JSON(
			http.StatusOK,
			res,
		)
	}
}

func Pong(c *gin.Context) {
	pong := model.PingResponse{Message: "pong"}
	c.JSON(
		http.StatusOK,
		pong,
	)
}

func handlerContext() context.Context {
	return context.Background()
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

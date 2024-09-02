package dao

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/blinlol/bookworm/model"
)


var logger *zap.Logger
var connString string

// return nil in was error
func AllBooks(ctx context.Context) []*model.Book {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		logger.Sugar().Error(err)
		return nil
	}

	var id, title, author string
	books := make([]*model.Book, 0)
	rows, _ := conn.Query(ctx, "select id, title, author from books")
	defer rows.Close()
	_, err = pgx.ForEachRow(
		rows,
		[]any{&id, &title, &author},
		func() error {
			b := model.Book{Id: id, Title: title, Author: author}
			books = append(books, &b)
			return nil
		},
	)

	if err != nil {
		logger.Sugar().Error(err)
		return nil
	}

	rows.Close()
	if err = rows.Err(); err != nil {
		logger.Sugar().Error(err)
		return nil
	}
	return books
}


func FindBookById(ctx context.Context, id string) *model.Book {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		logger.Sugar().Error(err)
		return nil
	}

	book := model.Book{Id: id}
	err = conn.QueryRow(
		ctx,
		"select author, title from books where id = $1",
		id,
	).Scan(&book.Author, &book.Title)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}
		logger.Sugar().Error(err)
		return nil
	}

	return &book
}


func FindBook(ctx context.Context, bookLike model.Book) *model.Book {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		logger.Sugar().Error(err)
		return nil
	}

	err = conn.QueryRow(
		ctx,
		"select id from books where author = $1 and title = $2",
		bookLike.Author, bookLike.Title,
	).Scan(&bookLike.Id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}
		logger.Sugar().Error(err)
		return nil
	}

	return &bookLike
}

// return nil if was error
func AddBook(ctx context.Context, book model.Book) *model.Book {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		logger.Sugar().Error(err)
		return nil
	}

	commandTag, err := conn.Exec(
		ctx,
		"insert into books(author, title) values ($1, $2)",
		book.Author, book.Title,
	)
	if err != nil {
		logger.Sugar().Error(err)
		return nil
	}
	if commandTag.RowsAffected() != 1 {
		logger.Sugar().Warn("insert affected more than 1 row, there is smth wrong")
	}

	return FindBook(ctx, book)
}

func DeleteBookById(ctx context.Context, id string) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		logger.Sugar().Error(err)
		return
	}

	commandTag, err := conn.Exec(
		ctx,
		"delete from books where id = $1",
		id,
	)
	if err != nil {
		logger.Sugar().Error(err)
	} else if commandTag.RowsAffected() == 0 {
		logger.Sugar().Infof("book with id = %s not found and cant delete", id)
	}
}

// func UpdateBook(book *model.Book) *model.Book {}


func init(){
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	connString = "postgresql://bookworm_user:123@localhost:5432/bookworm_db"
}
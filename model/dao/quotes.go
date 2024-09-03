package dao

import (
	"context"

	"github.com/blinlol/bookworm/model"
	"github.com/jackc/pgx/v5"
)


// return nil on error
func GetQuotesByBookId(ctx context.Context, id string) []*model.Quote {
	conn, err := pgx.Connect(
		ctx,
		ConnString,
	)
	if err != nil {
		Logger.Sugar().Error(err)
		return nil
	}

	rows, _ := conn.Query(
		ctx,
		"select quotes from quotes where book_id = $1",
		id,
	)
	
}
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

	var quotes_text []string
	err = conn.QueryRow(
		ctx,
		"select quotes from quotes where book_id = $1",
		id,
	).Scan(&quotes_text)

	if err != nil {
		if err == pgx.ErrNoRows {
			Logger.Sugar().Infof("quotes not found for id = %s", id)
		} else {
			Logger.Sugar().Error(err)
		}
		return nil
	}

	quotes := make([]*model.Quote, 0)
	for _, text := range quotes_text {
		quotes = append(quotes, &model.Quote{Text: text})
	}

	return quotes
}

func AddQuotes(ctx context.Context, bookId string, quotes []*model.Quote) (success bool) {
	conn, err := pgx.Connect(ctx, ConnString)
	if err != nil {
		Logger.Sugar().Error(err)
		return
	}

	quotes_for_query := make([]string, 0)
	for _, quote := range quotes {
		quotes_for_query = append(quotes_for_query, quote.Text)
	}
	commandTag, err := conn.Exec(
		ctx,
		"insert into quotes(book_id, quotes) values ($1, $2)",
		bookId, quotes_for_query,
	)
	if err != nil {
		Logger.Sugar().Error(err)
		return
	}
	if commandTag.RowsAffected() == 0 {
		Logger.Sugar().Warn("row not inserted")
		return
	}
	success = true
	return
}

func DeleteQuotes(ctx context.Context, bookId string) {
	conn, err := pgx.Connect(ctx, ConnString)
	if err != nil {
		Logger.Sugar().Error(err)
		return
	}

	_, err = conn.Exec(
		ctx,
		"delete from quotes where book_id = $1",
		bookId,
	)

	if err != nil {
		Logger.Sugar().Error(err)
	}
}

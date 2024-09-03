package dao

import (
	"context"
	"testing"

	"github.com/blinlol/bookworm/model"
	"github.com/stretchr/testify/assert"
)

func TestAddGetDelete(t *testing.T) {
	ctx := context.Background()
	b := &model.Book{Title: "title", Author: "author"}
	b = AddBook(ctx, *b)
	quotes := []*model.Quote{
		{Text: "pupupu"},
		{Text: "papapa"},
	}
	AddQuotes(ctx, b.Id, quotes)
	quotes_from_get := GetQuotesByBookId(ctx, b.Id)
	assert.Equal(t, quotes, quotes_from_get)
	DeleteQuotes(ctx, b.Id)
	DeleteBookById(ctx, b.Id)
}

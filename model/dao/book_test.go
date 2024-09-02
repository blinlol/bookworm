package dao

import (
	"context"
	"testing"
	"github.com/blinlol/bookworm/model"
	"github.com/stretchr/testify/assert"
)


func TestAddFindDeleteFind(t *testing.T) {
	ctx := context.Background()
	new_b := &model.Book{Title: "A", Author: "B"}
	from_add_b := AddBook(ctx, *new_b)
	assert.Equal(t, new_b.Title, from_add_b.Title)
	assert.Equal(t, new_b.Author, from_add_b.Author)
	b := FindBook(ctx, *new_b)
	assert.Equal(t, *from_add_b, *b)

	DeleteBookById(ctx, b.Id)
	b = FindBook(ctx, *new_b)
	assert.Nil(t, b)
}

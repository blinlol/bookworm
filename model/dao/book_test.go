package dao

import (
	"context"
	"testing"
	"github.com/blinlol/bookworm/model"
	"github.com/stretchr/testify/assert"
)


func TestAddFindUpdateDeleteFind(t *testing.T) {
	ctx := context.Background()
	new_b := &model.Book{Title: "A", Author: "B"}
	// add
	from_add_b := AddBook(ctx, *new_b)
	assert.Equal(t, new_b.Title, from_add_b.Title)
	assert.Equal(t, new_b.Author, from_add_b.Author)
	// find
	b := FindBook(ctx, *new_b)
	assert.Equal(t, *from_add_b, *b)
	//update
	b.Author = "1"
	b.Title = "2"
	UpdateBook(ctx, *b)
	found_b := FindBookById(ctx, b.Id)
	assert.Equal(t, *b, *found_b)
	//delete
	DeleteBookById(ctx, b.Id)
	found_b = FindBook(ctx, *b)
	assert.Nil(t, found_b)
}

package dao

import (
	"testing"

	"github.com/blinlol/bookworm/model"
	"github.com/stretchr/testify/assert"
)

func TestAddFindDeleteFind(t *testing.T) {
	new_b := &model.Book{Title: "A", Author: "B"}
	AddBook(new_b)
	b := FindBook("A", "B")
	if b == nil {
		t.Error("not found book")
	} else if b.Author != "B" && b.Title != "A" {
		t.Error("wrong book")
	}
	DeleteBook(b)
	b = FindBook("A", "B")
	if b != nil {
		t.Errorf("book not deleted")
	}
}

func TestAddQuote(t *testing.T) {
	new_b := &model.Book{Title: "A", Author: "B"}
	AddBook(new_b)
	b := FindBook("A", "B")
	q := &model.Quote{Text: "qwerty"}
	AddQuote(b, q)
	b = FindBook("A", "B")
	if len(b.Quotes) != 1 || *(b.Quotes[0]) != *q {
		t.Error("quote not saved")
	}
	DeleteBook(new_b)
}

func TestFindAll(t *testing.T) {
	b1 := &model.Book{Title: "A", Author: "B"}
	b2 := &model.Book{Title: "C", Author: "D"}

	AddBook(b1)
	AddBook(b2)
	all := AllBooks()

	if len(all) != 2 {
		t.Error("wrong books count: ", len(all))
	}
	DeleteBook(b1)
	DeleteBook(b2)
}

func TestUpdate(t *testing.T) {
	b := &model.Book{Title: "1", Author: "2"}
	b = AddBook(b)

	b.Author = "a"
	b.Title = "t"
	updated_b := UpdateBook(b)
	assert.Equal(t, *b, *updated_b)

	DeleteBook(b)
}

package dao

import (
	"testing"
)


func TestAddFindDeleteFind(t *testing.T){
	AddBook("A", "B")
	b := FindBook("A", "B")
	if b == nil {
		t.Error("not found book")
	} else if b.Author != "B" && b.Title != "A" {
		t.Error("wrong book")
	}
	DeleteBook("A", "B")
	b = FindBook("A", "B")
	if b != nil {
		t.Errorf("book not deleted")
	}
}


func TestAddQuote(t *testing.T){
	AddBook("A", "B")
	b := FindBook("A", "B")
	q := &Quote{Text: "qwerty"}
	AddQuote(b, q)
	b = FindBook("A", "B")
	if len(b.Quotes) != 1 || *(b.Quotes[0]) != *q {
		t.Error("quote not saved")
	}
	DeleteBook("A", "B")
}


func TestFindAll(t *testing.T){
	AddBook("A", "B")
	AddBook("C", "D")
	all := AllBooks()

	if len(all) != 2 {
		t.Error("wrong books count")
	}
	DeleteBook("A", "B")
	DeleteBook("C", "D")
}
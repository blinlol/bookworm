package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blinlol/bookworm/model"
	"github.com/blinlol/bookworm/model/dao"
)


func TestPong(t *testing.T){
	r := CreateRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}


func TestBooks(t *testing.T){
	// init router
	r := CreateRouter()
	r = BookRoutes(r)

	// add book
	w := httptest.NewRecorder()
	b := model.AddBookRequest{Book: model.Book{Title: "title", Author: "author"}}
	raw, _ := json.Marshal(b)
	req, _ := http.NewRequest(
		"POST",
		"/api/books/add",
		bytes.NewReader(raw),
	)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var resp_b model.AddBookResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp_b)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, b.Book.Author, resp_b.Book.Author)
	assert.Equal(t, b.Book.Title, resp_b.Book.Title)

	// find existing book
	w = httptest.NewRecorder()
	id := resp_b.Book.Id
	req, _ = http.NewRequest(
		"GET",
		"/api/book/" + id,
		nil,
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var resp_b_2 model.GetBookResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp_b_2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp_b.Book, *resp_b_2.Book)

	// find not existing book
	w = httptest.NewRecorder()
	id = "not_found"
	req, _ = http.NewRequest(
		"GET",
		"/api/book/" + id,
		nil,
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp_b_3 model.ErrorResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp_b_3)
	if err != nil {
		t.Error(err)
	}

	// delete book
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(
		"DELETE",
		"/api/book/" + resp_b.Book.Id,
		nil,
	)
	r.ServeHTTP(w, req)

	assert.Nil(t, dao.FindBook("title", "author"))
}

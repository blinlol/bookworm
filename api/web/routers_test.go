package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blinlol/bookworm/data/dao"
)


func TestPong(t *testing.T){
	r := SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}


func TestBooks(t *testing.T){
	r := SetupRouter()
	r = BookRoutes(r)
	w := httptest.NewRecorder()

	b := dao.Book{Title: "title", Author: "author"}
	raw, _ := json.Marshal(b)
	req, _ := http.NewRequest(
		"POST",
		"/books/add",
		bytes.NewReader(raw),
	)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var resp_b dao.Book
	err := json.Unmarshal(w.Body.Bytes(), &resp_b)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, b.Author, resp_b.Author)
	assert.Equal(t, b.Title, resp_b.Title)

	req, _ = http.NewRequest(
		"DELETE",
		"/book/" + resp_b.Id,
		nil,
	)
	r.ServeHTTP(w, req)

	assert.Nil(t, dao.FindBook("title", "author"))
}
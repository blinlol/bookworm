package model

type Book struct {
	Id     string
	Title  string
	Author string
	Quotes []*Quote
}

type Quote struct {
	Text	string
}

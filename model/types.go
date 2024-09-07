package model

type Book struct {
	Id     string	`json:"id"`
	Title  string	`json:"title"`
	Author string	`json:"author"`
}

type Quote struct {
	Text	string	`json:"text"`
}

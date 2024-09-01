package model

type Book struct {
	// hex representation of ObjectID
	Id     string `bson:"_id"`
	Title  string
	Author string
	Quotes []*Quote
}

type Quote struct {
	Text	string
}

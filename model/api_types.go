package model

type ErrorResponse struct {
	Message	string	`json:"message"`
}

type GetBooksResponse struct {
	Books	[]*Book	`json:"books"`
}

type GetBookResponse struct {
	Book	*Book	`json:"book"`
}

type AddBookRequest struct {
	Book	Book	`json:"book"`
}

type AddBookResponse struct {
	Book	Book	`json:"book"`
}

type UpdateBookRequest struct {
	Book	Book	`json:"book"`
}

type ParseQuotesRequest struct {
	BookId 	string	`json:"book_id"`
	Text	string	`json:"text"`
	Separator	string	`json:"separator"`
}

type PingResponse struct {
	Message	string	`json:"message"`
}

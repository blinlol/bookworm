package model


type ErrorResponse struct {
	Message string	`json:"message"`
}


type GetBooksResponse struct {
	Books []*Book	`json:"books"`
}


type GetBookResponse struct {
	Book *Book `json:"book"`
}


type AddBookRequest struct {
	Book Book	`json:"book"`
}

type AddBookResponse struct {
	Book Book	`json:"book"`
}


type PingResponse struct {
	Message string	`json:"message"`
}

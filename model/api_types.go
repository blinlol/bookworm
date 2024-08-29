package model


type GetBooksRequest struct {
}

type GetBooksResponse struct {
	Books []Book	`json:"books"`
}


type AddBookRequest struct {
	Book Book	`json:"book"`
}

type AddBookResponse struct {
	Book Book	`json:"book"`
}


type PingResponse struct {
	Message string
}

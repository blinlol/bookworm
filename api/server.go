package main

import (
	"log"

	"github.com/blinlol/bookworm/api/web"
)

func main(){
	address := "localhost:8080"
	err := web.StartServer(address)
	if err != nil {
		log.Fatalln(err)
	}
}
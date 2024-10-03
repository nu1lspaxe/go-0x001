package main

import (
	"log"
	"net/http"

	"github.com/nu1lspaxe/go-0x001/server/swagger"
)

func main() {
	log.Printf("Server started")

	router := swagger.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

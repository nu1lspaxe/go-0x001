package main

import (
	"go_0x001/server/swagger"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := swagger.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

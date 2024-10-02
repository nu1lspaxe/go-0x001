package main

import (
	"log"
	"net/http"
	"server/swagger"
)

func main() {
	log.Printf("Server started")

	router := swagger.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

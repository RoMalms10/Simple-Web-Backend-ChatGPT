package main

import (
	"log"
	"net/http"

	"Simple-Web-Backend-ChatGPT/server"
)

func main() {
	r := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}

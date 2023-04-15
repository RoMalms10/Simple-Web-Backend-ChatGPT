package server

import (
	"net/http"

	"Simple-Web-Backend-ChatGPT/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/people", handler.GetPeopleHandler).Methods(http.MethodGet)
	r.HandleFunc("/people", handler.PostPeopleHandler).Methods(http.MethodPost)
	r.HandleFunc("/people", handler.DeletePeopleHandler).Methods(http.MethodDelete)

	return r
}

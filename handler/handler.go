package handler

import (
	"encoding/json"
	"net/http"

	"Simple-Web-Backend-ChatGPT/models"
)

func GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		// Get a specific person by name
		for _, p := range models.People {
			if p.Name == name {
				json.NewEncoder(w).Encode(p)
				return
			}
		}
		http.Error(w, "Person not found.", http.StatusNotFound)
		return
	}

	// Return the entire list of people
	json.NewEncoder(w).Encode(models.People)
}

func PostPeopleHandler(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.AddPerson(person)
	w.WriteHeader(http.StatusCreated)
}

func DeletePeopleHandler(w http.ResponseWriter, r *http.Request) {
	var nameToDelete string
	err := json.NewDecoder(r.Body).Decode(&nameToDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !models.DeletePersonByName(nameToDelete) {
		http.Error(w, "Person not found.", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

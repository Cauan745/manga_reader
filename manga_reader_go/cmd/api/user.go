package main

import (
	"encoding/json"
	"net/http"

	"manga-reader/internal/models"
)

func (a *application) registerUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid User json", 400)
	}

	w.WriteHeader(200)
	w.Write([]byte("User registered successfully!"))
}

package main

import (
	"encoding/json"
	"net/http"
	"time"

	repo "manga-reader/internal/adapters/postgresql/sqlc"
	"manga-reader/internal/services"

	"github.com/golang-jwt/jwt/v5"
)

func (a *application) registerUser(w http.ResponseWriter, r *http.Request) {
	secret := "my-secret"
	issuer := "my-server"

	user := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid User json", 400)
		return
	}

	hash, err := services.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// a.logger.Info("Creating password:", "Generated Hash", hash)

	userParams := repo.CreateUserParams{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: hash,
	}

	dbUser, err := a.db.CreateUser(r.Context(), userParams)
	if err != nil {
		// a.logger.Warn("Error saving to database", "params", userParams)
		http.Error(w, "This username or email is already u", http.StatusBadRequest)
		return
	}

	a.logger.Info("User created", "username", dbUser.Username)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": issuer,
			"sub": user.Email,
			"iat": time.Now().Unix(),
		},
	)

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		http.Error(w, "Error generating token", 500)
		return
	}

	a.logger.Info(token)

	type response struct {
		Token string `json:"token"`
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(response{token})
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

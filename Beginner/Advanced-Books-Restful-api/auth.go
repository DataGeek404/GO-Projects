package main

import (
	"encoding/json"
	"net/http"
)

var users = map[string]string{} // In-memory user storage

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	_ = json.NewDecoder(r.Body).Decode(&creds)

	if password, ok := users[creds.Username]; ok && password == creds.Password {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Invalid credentials"))
}

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	_ = json.NewDecoder(r.Body).Decode(&creds)

	if _, exists := users[creds.Username]; exists {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("User already exists"))
		return
	}

	users[creds.Username] = creds.Password
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered"))
}

// Credentials holds user login information
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

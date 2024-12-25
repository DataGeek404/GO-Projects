package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Middleware
	r.Use(loggingMiddleware)

	// Book routes
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	// Authentication routes
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/register", Register).Methods("POST")

	// Start server at point 8080
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Middleware to log incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// Book Handlers
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all books"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Get book with ID: " + vars["id"]))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new book"))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Update book with ID: " + vars["id"]))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Delete book with ID: " + vars["id"]))
}

// Authentication Handlers
func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login handler"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register handler"))
}

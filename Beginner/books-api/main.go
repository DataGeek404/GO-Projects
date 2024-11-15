package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

// Book represents a book object
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// In-memory book collection
var books []Book

// Handler to get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// Handler to get a single book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, book := range books {
        if book.ID == params["id"] {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler to create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(1000000)) // Generate a random ID
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

// Handler to update an existing book
func updateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for i, book := range books {
        if book.ID == params["id"] {
            books = append(books[:i], books[i+1:]...)
            var updatedBook Book
            _ = json.NewDecoder(r.Body).Decode(&updatedBook)
            updatedBook.ID = params["id"]
            books = append(books, updatedBook)
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler to delete a book by ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for i, book := range books {
        if book.ID == params["id"] {
            books = append(books[:i], books[i+1:]...)
            json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Main function to set up routes and start the server
func main() {
    router := mux.NewRouter()

    // Seed some initial data
    books = append(books, Book{ID: "1", Title: "The Catcher in the Rye", Author: "J.D. Salinger"})
    books = append(books, Book{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"})

    // Route Handlers / Endpoints
    router.HandleFunc("/api/books", getBooks).Methods("GET")
    router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    router.HandleFunc("/api/books", createBook).Methods("POST")
    router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
    router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

    // Start the server
    fmt.Println("Starting server on port 8000...")
    log.Fatal(http.ListenAndServe(":8000", router))
}

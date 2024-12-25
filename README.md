# Books API

A simple RESTful API for managing a collection of books. This API allows you to create, read, update, and delete (CRUD) books with basic data such as `ID`, `Title`, and `Author`. Built using Go and the Gorilla Mux router.

## Features

- **Create a new book** with title and author details.
- **Retrieve all books** or a single book by its ID.
- **Update a book** by ID to modify title and author.
- **Delete a book** by its ID.

## Requirements

- [Go](https://golang.org/dl/) installed on your system.
- [Gorilla Mux](https://github.com/gorilla/mux) package installed for routing.

To install Gorilla Mux, run:
```bash
go get -u github.com/gorilla/mux
```

## Getting Started
## Clone and Set Up the Project
Clone this repository to your local machine:

```bash
git clone https://github.com/yourusername/books-api.git
cd books-api
```
Initialize and set up the Go module:

```bash
go mod init books-api
go mod tidy
```
Run the server:

```bash
go run main.go
```
## API Endpoints

| Method | Endpoint            | Description                |
|--------|----------------------|----------------------------|
| GET    | `/api/books`        | Retrieve all books         |
| GET    | `/api/books/{id}`   | Retrieve a single book     |
| POST   | `/api/books`        | Create a new book          |
| PUT    | `/api/books/{id}`   | Update an existing book    |
| DELETE | `/api/books/{id}`   | Delete a book by its ID    |

## 1. Get All Books
Request:

```bash
curl -X GET http://localhost:8000/api/books
```
Response:
```bash
[
    {
        "id": "1",
        "title": "The Catcher in the Rye",
        "author": "J.D. Salinger"
    },
    {
        "id": "2",
        "title": "To Kill a Mockingbird",
        "author": "Harper Lee"
    }
]

```
## Get a Book by ID
Request:
```bash
curl -X GET http://localhost:8000/api/books/1
```

Response:

```bash
{
    "id": "1",
    "title": "The Catcher in the Rye",
    "author": "J.D. Salinger"
}

```

## Create a New Book
Request:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "1984", "author": "George Orwell"}' http://localhost:8000/api/books

```
Response:

```bash
{
    "id": "randomly_generated_id",
    "title": "1984",
    "author": "George Orwell"
}

```

## Update a Book
Request:
```bash 
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Animal Farm", "author": "George Orwell"}' http://localhost:8000/api/books/1
```
Response:
```bash
{
    "id": "1",
    "title": "Animal Farm",
    "author": "George Orwell"
}
```
## Delete a Book
Request:
```bash
curl -X DELETE http://localhost:8000/api/books/1

```
Response:

```bash
{
    "message": "Book deleted"
}

```
## Project Structure
main.go: The main application file containing all routes and handlers for CRUD operations.
Book struct: Defines the data structure for each book entry.
Handlers: Each CRUD operation (GET, POST, PUT, DELETE) has a dedicated function to handle requests and manage the in-memory book collection.
## Future Improvements
Database Integration: Replace the in-memory storage with a persistent database (e.g., PostgreSQL, MySQL).
Authentication: Add authentication to secure the API endpoints.
Pagination: Implement pagination for the GET /api/books endpoint when the book list grows.

### Additional Notes

This README provides all essential information:
- Setup instructions
- API endpoints and example requests/responses
- Project structure and potential future improvements

Let me know if you'd like to add any more details or further customization!

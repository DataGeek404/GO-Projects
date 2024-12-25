# Advanced Books API

An advanced RESTful API built with Go for managing a collection of books. This API features CRUD operations, JWT-based authentication, request logging, and pagination.

## Features

- **Create, Read, Update, Delete** (CRUD) books.
- **JWT Authentication** for secure access.
- **Pagination** for GET requests.
- **Logging** of requests and response times.
- **PostgreSQL Database** for data persistence.

## Requirements

- [Go](https://golang.org/dl/)
- PostgreSQL
- [Gorilla Mux](https://github.com/gorilla/mux)
- [JWT-go](https://github.com/dgrijalva/jwt-go)
- [Logrus](https://github.com/sirupsen/logrus)

## Setup

1. Clone the repository and navigate into the project directory.
2. Set up PostgreSQL and create a `books` table.
3. Configure your PostgreSQL credentials in `models.go`.
4. Install dependencies:
   ```bash
   go mod tidy
## Run the server:

```bash
go run main.go


```
# API Endpoints

This document outlines the available API endpoints for managing books in the system.

| Method | Endpoint            | Description                          |
|--------|---------------------|--------------------------------------|
| POST   | /login              | Generate JWT token                  |
| GET    | /api/books          | Get all books (paginated)            |
| POST   | /api/books          | Create a new book                   |
| PUT    | /api/books/{id}     | Update an existing book             |
| DELETE | /api/books/{id}     | Delete a book by ID                 |

## Authentication

- The `/login` endpoint is used to generate a JWT token for authenticating requests.

## Book Management

- The `/api/books` endpoint allows retrieving all books, creating a new book, and updating or deleting existing books by specifying the book ID in the URL.

## License
## MIT License
```bash

This README provides essential details, setup instructions, and API documentation for your advanced RESTful API project in Go. Let me know if you’d like more details on any aspect!
```

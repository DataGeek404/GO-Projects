# Simple Web Server in Go

This project is a basic web server written in Go. When accessed, it responds with a custom message. It’s a simple project to demonstrate the basics of setting up a server in Go.

## Project Details

- **Message**: Responds with `"My name is James Muchiri, This is my first server, brave!!"` at the root URL (`"/"`).
- **Port**: Runs on port `8080` by default.

## Features

- Basic HTTP server setup using Go’s `net/http` package.
- Custom message handling for requests to the root URL.
- Error handling for server startup issues.

## Requirements

- [Go](https://golang.org/dl/) must be installed on your system.

## Getting Started

1. **Clone or download** this project.
2. **Navigate** to the directory containing the `main.go` file.
3. Run the server with the following command:

   ```bash
   go run main.go


## You should see the message:

```bash
Starting server on port 8080...
```
## Test the Server:
Open a web browser and go to http://localhost:8080.
You should see the message: My name is James Muchiri, This is my first server, brave!!

## Explanation of the Code

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {...}):
Sets up a handler for requests to the root URL ("/").
When accessed, it responds with a custom message.
Starting the Server:
http.ListenAndServe(":8080", nil) starts the HTTP server on port 8080.
If an error occurs (e.g., port 8080 is in use), it will print the error to the console.

## Example Output

```bash 
Starting server on port 8080...

```
## In your browser at http://localhost:8080:

```bash
My name is James Muchiri, This is my first server, brave!!
```

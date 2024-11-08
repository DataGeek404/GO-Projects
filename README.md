Creating a Simple "Hello, World!" Web Server in Go

1. Set Up the Project:
- Create a new directory for your project and navigate into it:

mkdir hello-world-server

cd hello-world-server

- Initialize a new Go module:

go mod init hello-world-server

2\. Write the Code:

- Create a file named main.go and add the following content:

package main

import (

"fmt"

"net/http"

)

func handler(w http.ResponseWriter, r \*http.Request) {

fmt.Fprintln(w, "Hello, World!")

}

func main() {

http.HandleFunc("/", handler)

fmt.Println("Starting server on :8080")

if err := http.ListenAndServe(":8080", nil); err != nil {

fmt.Println("Error starting server:", err)

}

}

3\. Run the Server:

- In your terminal, execute:

go run main.go

- You should see the message:

Starting server on :8080

4\. Test the Server:

- Open a web browser and navigate to http://localhost:8080. You should see the message "Hello,

World!" displayed.

Explanation:

- The main package is the entry point for the Go application.
- The net/http package provides HTTP client and server implementations.
- The handler function writes "Hello, World!" to the HTTP response.
- http.HandleFunc("/", handler) registers the handler function to handle requests to the root URL

path ("/").

- http.ListenAndServe(":8080", nil) starts the HTTP server on port 8080. The nil parameter indicates

that the DefaultServeMux will be used, which is where HandleFunc registers the handler.

This basic web server listens on port 8080 and responds with "Hello, World!" to any HTTP requests

made to the root URL.

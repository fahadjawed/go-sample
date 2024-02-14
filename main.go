package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function
    handler := func(w http.ResponseWriter, r *http.Request) {
        // Write response to the client
        fmt.Fprintf(w, "Hello, this is a basic Go lang server!")
    }

    // Register the handler function to handle all requests at "/"
    http.HandleFunc("/", handler)

    // Start the HTTP server on port 3000
    fmt.Println("Server is listening on port 3000...")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
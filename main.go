package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
    handleRoot := func(w http.ResponseWriter, _ *http.Request) {
        io.WriteString(w, "Welcome to simple hello world app!\n Go to /hello")
    }
    handleHello := func(w http.ResponseWriter, _ *http.Request) {
        io.WriteString(w, "Hello, World!")
    }
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/hello", handleHello)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

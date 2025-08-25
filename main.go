// Chapter 2.2: Web Application Basics
// Added a basic web server with a single route "/" handled by the home function.
// This step introduces http.NewServeMux, route handling, and starting the server on port 4000.

package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From SnippetBox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Server Listening on Port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

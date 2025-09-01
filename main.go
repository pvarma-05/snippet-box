package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From SnippetBox"))
}

func snippetView(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Displaying specific Snippet"))
}

func snippetCreate(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Displaying a form for creating a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Server Listening on Port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

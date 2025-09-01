package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.Atoi(r.PathValue("id"))
	if err!= nil || id<1{
		http.NotFound(w,r)
		return
	}
	msg := fmt.Sprintf("Displaying a specific snippet with ID: %d",id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying a form for creating a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Server Listening on Port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

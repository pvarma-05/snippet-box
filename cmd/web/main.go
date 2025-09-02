package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr",":4000","HTTP Network Address")
	flag.Parse()
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/",http.StripPrefix("/static",fs))
	mux.HandleFunc("GET /{$}",home)
	mux.HandleFunc("GET /snippet/view/{id}",snippetView)
	mux.HandleFunc("GET /snippet/create",snippetCreate)
	mux.HandleFunc("POST /snippet/create",snippetCreatePost)

	log.Printf("Server Started on %s",*addr)
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}

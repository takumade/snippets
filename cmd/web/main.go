package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /coffee/{id}", coffeeView)
	mux.HandleFunc("POST /coffee/add", coffeeAdd)
	mux.HandleFunc("DELETE /coffee/delete/{id}", coffeeDelete)

	log.Print("Starting server on :5000")

	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
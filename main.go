package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from coffee land!"))
}

func coffeeView(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Server", "Napkins")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("View coffee!"))

	
}

func coffeeDelete(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("View delete!"))

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	
	msg := fmt.Sprintf("Delete a specific coffee with ID %d...", id)
    w.Write([]byte(msg))

	
}

func coffeeAdd(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Add coffee!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /coffee/{id}", coffeeView)
	mux.HandleFunc("POST /coffee/add", coffeeAdd)
	mux.HandleFunc("DELETE /coffee/delete/{id}", coffeeDelete)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
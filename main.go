package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from coffee land!"))
}

func coffeeView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("View coffee!"))
}

func coffeeDelete(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("View delete!"))
}

func coffeeAdd(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("View delete!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/coffee/view", coffeeView)
	mux.HandleFunc("/coffee/add", coffeeAdd)
	mux.HandleFunc("/coffee/delete", coffeeDelete)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
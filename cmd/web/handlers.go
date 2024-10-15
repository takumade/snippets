package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)




func home(w http.ResponseWriter, r *http.Request){

	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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

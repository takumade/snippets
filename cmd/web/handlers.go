package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")
	// Include the navigation partial in the template files.
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Napkins")
	w.WriteHeader(http.StatusAccepted)

	id := r.PathValue("id")
	w.Write([]byte("View a specific coffee with ID " + id + "..."))

}

func (app *application) snippetDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View delete!"))

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Delete a specific coffee with ID %d...", id)
	w.Write([]byte(msg))

}

func (app *application) snippetAdd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add coffee!"))
}

func (app *application) snippetCreatePost (w http.ResponseWriter, r *http.Request){
	title := "O snail"
	content := "Climb Mount Fuji, But slowly, slowly!"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

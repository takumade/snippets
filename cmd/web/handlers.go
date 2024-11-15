package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"snippetbox.takucoder.dev/internal/models"
	"snippetbox.takucoder.dev/internal/validator"
)

type snippetCreateForm struct {
	Title string 
	Content string 
	Expires int
	validator.Validator
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")
	// Include the navigation partial in the template files.

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.html", data)

	
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w,r,http.StatusOK, "view.html", data)
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

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request){
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create.html", data)
}

func (app *application) snippetAdd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add coffee!"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {


	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return 
	}

	form := snippetCreateForm{
		Title: r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: expires,
	}

	// Validation 
	form.CheckFiled(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckFiled(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckFiled(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckFiled(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")


	if len(form.FieldErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
		return
	}


	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}


func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData){
	ts, ok := app.templateCache[page]

	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return 
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
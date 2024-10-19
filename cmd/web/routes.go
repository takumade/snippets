package main
import "net/http"
// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/{id}", app.snippetView)
	mux.HandleFunc("POST /snippet/add", app.snippetAdd)
	mux.HandleFunc("DELETE /snippet/delete/{id}", app.snippetDelete)
    return mux
}
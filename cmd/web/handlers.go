
package  main 


import (
	"fmt"
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

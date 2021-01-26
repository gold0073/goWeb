package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get User in for by /users/{id}")
}

func getUserInforHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "User id:", vars["id"])
}

//NewHandler is
func NewHandler() http.Handler {
	mux := mux.NewRouter()
	//mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInforHandler)
	return mux
}

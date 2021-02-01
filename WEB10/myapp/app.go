package myapp

import (
	"fmt"
	"net/http"
)

//IndexHandler is ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

//NewHandler is...
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	return mux
}

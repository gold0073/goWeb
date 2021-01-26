package myapp

import (
	"net/http"

	"github.com/gorilla/mux"
)


func NewHnadler() http.Handler
{
	mux : = http.NewServeMux()
	return mux
}
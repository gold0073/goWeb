package main

import (
	"net/http"

	"github.com/gold0073/goWeb/web5/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}

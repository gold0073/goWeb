package main

import (
	"net/http"

	"goWeb/WEB5/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}

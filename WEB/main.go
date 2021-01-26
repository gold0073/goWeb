package main

import (
	"goWeb/WEB/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHTTPHandler())

}

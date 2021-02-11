package main

import (
	"goWeb/WEB20/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	filepath := "./test.db"
	m := app.MakeHandler(filepath)
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)

	if err != nil {
		panic(err)
	}

}

package main

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func main() {
	fmt.Println(os.Getenv("SESSION_KEY"))
}

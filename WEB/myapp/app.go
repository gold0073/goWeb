package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//User is ...
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request:", err)
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w, string(data))

	//fmt.Fprint(w, "Hello Foo")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "barPark"
	}

	fmt.Fprintf(w, "Hello %s", name)
}

//NewHTTPHandler is...
func NewHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	return mux
}

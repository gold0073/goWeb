package app

import (
	"net/http"
	"strconv"

	"goWeb/WEB20/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DbHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.db.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func MakeHandler(filepath string) *AppHandler {
	//todoMap = make(map[int]*Todo)

	r := mux.NewRouter()
	app := &AppHandler{
		Handler: r,
		db:      model.NewDbHandler(filepath),
	}

	r.HandleFunc("/todos", app.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", app.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", app.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", app.completeTodoHandler).Methods("GET")
	r.HandleFunc("/", app.indexHandler)

	return app
}

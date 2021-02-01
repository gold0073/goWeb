package main

import (
	"goWeb/WEB10/decoHandler"
	"goweb/web10/myapp"
	"log"
	"net/http"
	"time"
)

//logger is ...
func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	startTime := time.Now()
	log.Println("[log1] Started")
	h.ServeHTTP(w, r)
	log.Println("[log1] Completed Time:", time.Since(startTime).Microseconds())
}

//logger2 is ...
func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	startTime := time.Now()
	log.Println("[log2] Started")
	h.ServeHTTP(w, r)
	log.Println("[log2] Completed Time:", time.Since(startTime).Microseconds())
}

//NewHandler is...
func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decoHandler.NewDecoHandler(h, logger)
	h = decoHandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)

}

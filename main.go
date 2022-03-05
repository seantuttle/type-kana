package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("./templates/*"))

func main() {
	logFile, err := os.OpenFile("./tmp/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(logging(static()))
	r.Handle("/", logging(index()))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}

	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("main: running type-kana server on port", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: couldn't start simple server: %v\n", err)
	}
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL, "handled in", time.Since(start))
	})
}

// index is the handler responsible for rending the index page for the site.
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "index", nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
	})
}

func static() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
}

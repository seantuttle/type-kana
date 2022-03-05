package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates = template.Must(template.ParseGlob("../templates/*"))

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL, "handled in", time.Since(start))
	})
}

// index is the handler responsible for rending the index page for the site.
func Index() http.Handler {
	return logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "index", nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
	}))
}

func Static() http.Handler {
	return logging(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

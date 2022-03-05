package api

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := template.Must(template.ParseGlob("../templates/*")).ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
		return
	}
}

// func Static() http.Handler {
// 	return logging(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
// }

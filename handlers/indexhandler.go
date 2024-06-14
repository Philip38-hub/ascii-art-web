package handler

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

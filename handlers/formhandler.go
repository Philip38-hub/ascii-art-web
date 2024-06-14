package handler

import (
	asciiart "asciiart/functions"
	"html/template"
	"net/http"
)

type AsciiData struct {
	Text   string
	Banner string
	Result string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "error 405: Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	inputString := r.FormValue("text")
	banner := r.FormValue("banner")

	data := AsciiData{
		Text:   inputString,
		Banner: banner,
		Result: asciiart.CreateCharacters(banner, inputString),
	}

	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, "error 500: internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "error 500: internal server error", http.StatusInternalServerError)
	}
}

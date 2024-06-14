package main

import (
	handler "asciiart/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoute)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": handler.IndexHandler(w)
	case "/ascii-art": handler.FormHandler(w, r)
	default:
		http.Error(w, "error 404: page not found", http.StatusNotFound)
	}
}
package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	asciiart "asciiart/main"
)

type AsciiData struct {
	Text   string
	Banner string
	Result string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	inputString := r.FormValue("text")
	banner := r.FormValue("banner")

	data := AsciiData{
		Text:   inputString,
		Banner: banner,
		Result: createCharacters(banner, inputString),
	}

	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", formHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createCharacters(bannername string, input string) string {
	file, _ := os.Open(bannername + ".txt")
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// group lines read into 9 to display each character individually
	character := make([]string, 9)
	var result [][]string
	for i := 1; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		character = lines[i:end]
		result = append(result, character)
	}

	// r := os.Args[1]
	// Format ("/n") in input string
	s := strings.Replace(input, "\r\n", "\n", -1)
	s = strings.Replace(s, "\t", "    ", -1)

	return asciiart.HandleLn(s, result)
}

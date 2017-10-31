package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlePage)
	http.HandleFunc("/api", handleAPI)
	if err := http.ListenAndServe(":10080", nil); err != nil {
		log.Fatal(err)
	}
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./main.html"))
	templ.Execute(w, nil)
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
	}
}

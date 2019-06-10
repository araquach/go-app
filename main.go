package main

import (
	"fmt"
	"google.golang.org/appengine"
	"html/template"
	"net/http"
)

type templateParams struct {
	Notice string
	Name string
}

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
	aboutTemplate = template.Must(template.ParseFiles("about.html"))
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	params := templateParams{}

	if r.Method =="GET" {
		indexTemplate.Execute(w, params)
		return
	}

	name := r.FormValue("name")
	params.Name = name
	if name == "" {
		name = "Anonymous Gopher"
	}

	if r.FormValue("message") == "" {
		w.WriteHeader(http.StatusBadRequest)

		params.Notice = "No Message provided"
		indexTemplate.Execute(w, params)
		return
	}

	params.Notice = fmt.Sprintf("Thanks you for your submission, %s", name)

	indexTemplate.Execute(w, params)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate.Execute(w, nil)
}
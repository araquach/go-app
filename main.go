package main

import (
	"google.golang.org/appengine"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template
var contactTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	// The path "/" matches everything not matched by some other path.
	http.Handle("/", r)
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml")
	if err != nil {
		panic(err)
	}
	appengine.Main() // Starts the server to receive requests
}
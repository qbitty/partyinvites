package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

type formData struct {
	*Rsvp
	Errors []string
}

var responses = make([]*Rsvp, 0, 10)

var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err != nil {
			panic(err)
		}

		templates[name] = t
		fmt.Println("Loaded template", index, name)
	}
}

func main() {
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

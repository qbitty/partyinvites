package main

import (
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templates["welcome"].Execute(w, nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	templates["list"].Execute(w, responses)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates["form"].Execute(w, formData{
			Rsvp:   &Rsvp{},
			Errors: []string{},
		})
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		responseData := Rsvp{
			Name:       r.Form["name"][0],
			Email:      r.Form["email"][0],
			Phone:      r.Form["phone"][0],
			WillAttend: r.Form["willattend"][0] == "true",
		}
		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email address")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}
		if len(errors) > 0 {
			templates["form"].Execute(w, formData{
				Rsvp: &responseData, Errors: errors,
			})
		} else {
			responses = append(responses, &responseData)

			if responseData.WillAttend {
				templates["thanks"].Execute(w, responseData.Name)
			} else {
				templates["sorry"].Execute(w, responseData.Name)
			}
		}
	}
}

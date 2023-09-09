package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "home.page.html")
}

// About func
func About(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "about.page.html")
}

func renderTemplate(response http.ResponseWriter, templ string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+templ, "./templates/base.layout.html")

	if err != nil {
		log.Println(err)

		return
	}

	err = parsedTemplate.Execute(response, nil)

	if err != nil {
		log.Println(err)

		return
	}
}

package render

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(response http.ResponseWriter, templ string) {
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

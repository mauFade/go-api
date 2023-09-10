package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplateDepr(response http.ResponseWriter, templ string) {
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

func RenderTemplate(response http.ResponseWriter, temp string) {
	var template *template.Template
	var err error

	_, isInCache := templateCache[temp]

	if !isInCache {
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(temp)

		if err != nil {
			log.Println(err)

			return
		}
	} else {
		log.Println("Using cache template")
	}

	template = templateCache[temp]

	err = template.Execute(response, nil)

	if err != nil {
		log.Println(err)

		return
	}
}

func createTemplateCache(temp string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", temp),
		"./templates/base.layout.html",
	}

	templ, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	templateCache[temp] = templ

	return nil
}

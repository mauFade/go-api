package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(response http.ResponseWriter, templ string) {
	templatesCache, err := createTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	temp, ok := templatesCache[templ]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(response)

	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templs, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			templs, err = templs.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return cache, err
			}
		}

		cache[name] = templs
	}

	return cache, nil
}

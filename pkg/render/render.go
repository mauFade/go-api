package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mauFade/go-api/pkg/config"
)

var templateCache = make(map[string]*template.Template)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(response http.ResponseWriter, templ string) {
	templatesCache, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	temp, ok := templatesCache[templ]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = temp.Execute(buf, nil)

	_, err = buf.WriteTo(response)

	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
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

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mauFade/go-api/pkg/config"
	"github.com/mauFade/go-api/pkg/handler"
	"github.com/mauFade/go-api/pkg/render"
)

const PORT = ":8081"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println("Application running on port", PORT)
	http.ListenAndServe(PORT, nil)
}

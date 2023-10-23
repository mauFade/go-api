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
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Application running on port", PORT)

	serve := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()

	log.Fatal(err)
}

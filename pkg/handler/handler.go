package handler

import (
	"net/http"

	"github.com/mauFade/go-api/pkg/render"
)

func Home(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, "home.page.html")
}

// About func
func About(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, "about.page.html")
}

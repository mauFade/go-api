package handler

import (
	"net/http"

	"github.com/mauFade/go-api/pkg/config"
	"github.com/mauFade/go-api/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, "home.page.html")
}

func (repo *Repository) About(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, "about.page.html")
}

package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/mauFade/go-api/pkg/config"
	"github.com/mauFade/go-api/pkg/handler"
)

func routes(*config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	return mux
}

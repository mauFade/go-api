package main

import (
	"net/http"

	"github.com/mauFade/go-api/pkg/handler"
)

const PORT = ":8081"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	http.ListenAndServe(PORT, nil)
}

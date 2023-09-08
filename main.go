package main

import (
	"net/http"

	"github.com/mauFade/go-api/handlers"
)

const PORT = ":8081"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(PORT, nil)
}

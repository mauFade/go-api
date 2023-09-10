package main

import (
	"fmt"
	"net/http"

	"github.com/mauFade/go-api/pkg/handler"
)

const PORT = ":8081"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println("Application running on port", PORT)
	http.ListenAndServe(PORT, nil)
}

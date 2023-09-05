package main

import (
	"errors"
	"fmt"
	"net/http"
)

const PORT = ":8081"

func Home(response http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(response, "Hello home page")

	if err != nil {
		return
	}
}

func About(response http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(response, "Hello about page")

	if err != nil {
		return
	}
}

func Divide(response http.ResponseWriter, request *http.Request) {
	float, err := divideValues(100.0, 0)

	if err != nil {
		fmt.Fprintf(response, err.Error())

		return
	}

	fmt.Fprintf(response, fmt.Sprintf("Valor da divisão: %f", float))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("A value cannot be divided by 0.")

		return 0, err
	}

	division := x / y

	return division, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	http.ListenAndServe(PORT, nil)
}

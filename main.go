package main

import (
	"net/http"

	response "./pkg"
)

func main() {

	http.HandleFunc("/", response.ResponseHandler)
	http.ListenAndServe(":8080", nil)
}

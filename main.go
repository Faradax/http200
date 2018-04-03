package main

import (
	"net/http"
)

var path = "/"

func allHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc(path, allHandler)
	http.ListenAndServe(":8080", nil)
}
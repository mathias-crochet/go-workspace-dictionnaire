package main

import (
	"dictionnaire/dictionary"
	"net/http"
)

func main() {

	dictionary.NewFile("dictionary.json")

	http.HandleFunc("/add", dictionary.Add)
	http.HandleFunc("/list", dictionary.List)
	http.HandleFunc("/get", dictionary.Get)
	http.HandleFunc("/remove", dictionary.Remove)

	http.ListenAndServe(":8080", nil)


}

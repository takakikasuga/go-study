package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello World111!")
}

// @desc net/http/server
func main() {
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServe(":8080", &MyHandler{})
}

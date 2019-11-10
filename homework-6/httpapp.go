package main

import (
	"net/http"
	"html/template"
	"log"
)

func hello(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("app/templates/template.html")

	if err != nil {
		log.Fatal(err)
	}

	name := req.URL.Query().Get("name")
	if name == "" {
		name = "Дружище"
	}
	data := struct {
		Name string
	}{
		Name: name,
	}
//log.Fatal(data)
	tmpl.Execute(res, data)
}

func main() {
	fs := http.FileServer(http.Dir("app/static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}

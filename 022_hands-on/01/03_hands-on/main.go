package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func serveTemplate(w http.ResponseWriter, s string)  {
	tpl.ExecuteTemplate(w, "results.gohtml", s)
}

func root(w http.ResponseWriter, req *http.Request)  {
	serveTemplate(w, "you are at root")
}

func dog(w http.ResponseWriter, req *http.Request)  {
	serveTemplate(w, "you are at dog")
}

func me(w http.ResponseWriter, req *http.Request)  {
	serveTemplate(w, "my name is go-lang-for-me")
}

func main()  {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
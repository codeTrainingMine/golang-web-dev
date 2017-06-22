package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main()  {
	http.HandleFunc("/", serveRoot)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)
}

func serveRoot(w http.ResponseWriter, req *http.Request)  {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
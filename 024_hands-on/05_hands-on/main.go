package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main()  {
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", serveRoot)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)
}

func serveRoot(w http.ResponseWriter, req *http.Request)  {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
package main

import (
	"net/http"
	"io"
	"html/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main()  {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dogpic/", dogpic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request)  {
	tpl.ExecuteTemplate(w, "dog.gohtml", "This is from dog")
}

func dogpic(w http.ResponseWriter, req *http.Request)  {
	http.ServeFile(w, req, "./assets/dog.jpeg")
}

package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request)  {
	c, err := req.Cookie("uid")
	if err != nil {
		c = &http.Cookie{Name: "uid", Value:uuid.NewV1().String()}
	} else {
		fmt.Println(c.Name, c.Value)
	}
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
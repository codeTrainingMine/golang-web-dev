package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template
var sunset string = "sunset.jpg"
var disney string = "disney.jpg"
var mickey string = "mickey.jpg"
var pipe string = "|"

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	vals := strings.Split(c.Value, pipe)
	tpl.ExecuteTemplate(w, "index.gohtml", vals)
}

// add func to get cookie
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	if !strings.Contains(c.Value, sunset) {
		c.Value += pipe + sunset
	}
	if !strings.Contains(c.Value, disney) {
		c.Value += pipe + disney
	}
	if !strings.Contains(c.Value, mickey) {
		c.Value += pipe + mickey
	}

	return c
}

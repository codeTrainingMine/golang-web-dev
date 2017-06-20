package main

import (
	"net/http"
	"io"
)

func root(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "you are at root")
}

func dog(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "you are at dog")
}

func me(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "my name is go-lang-for-me")
}

func main()  {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
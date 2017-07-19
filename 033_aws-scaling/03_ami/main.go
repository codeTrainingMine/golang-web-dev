package main

import (
	"net/http"
	"io"
)

func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Hello from AWS.")
}

func ping(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "instance two")
}


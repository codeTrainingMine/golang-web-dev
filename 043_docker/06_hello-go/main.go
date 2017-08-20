package main

import (
	"net/http"
	"io"
	"log"
)

func main()  {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello from a docker container!")
}
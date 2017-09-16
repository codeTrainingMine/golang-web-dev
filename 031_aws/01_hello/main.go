package main

import (
	"net/http"
	"io"
	"log"
)

func main()  {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Oh yeah, I'm running on the cloud.")
}
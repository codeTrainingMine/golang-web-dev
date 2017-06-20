package main

import (
	"fmt"
	"net/http"
	"log"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main()  {
	var d hotdog
	err := http.ListenAndServe(":80", d)
	if err != nil {
		log.Fatalln(err)
	}
}
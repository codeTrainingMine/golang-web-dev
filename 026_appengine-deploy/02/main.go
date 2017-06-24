package main

import "net/http"

func init()  {
	// test
	http.Handle("/", http.FileServer(http.Dir(".")))
}
package main

import (
	"net/http"
	"log"
	"io"
)

func main()  {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// the below will serve all files in /home/nir/
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

func index(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Oh yeah, I'm running on the cloud.")
}
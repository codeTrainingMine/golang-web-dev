package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func main()  {
	r := httprouter.New()
	r.GET("/", index)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "Welcome!\n")
}
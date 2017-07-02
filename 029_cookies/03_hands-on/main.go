package main

import (
	"net/http"
	"fmt"
	"strconv"
)

func main()  {
	http.HandleFunc("/", root)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request)  {
	var counter int
	c, err := req.Cookie("counter")
	if err != nil && err.Error() != "http: named cookie not present" {
		fmt.Println(err)
	} else {
		if err != nil && err.Error() == "http: named cookie not present" {
			counter = 1;
			writeCounter(w, counter)
		} else {
			var err2 error
			counter, err2 = strconv.Atoi(c.Value)
			if err2 != nil {
				fmt.Println(err2)
			} else {
				counter++
				writeCounter(w, counter)
			}
			//fmt.Fprintln(w, "test")
		}
	}
}

func writeCounter(w http.ResponseWriter, counter int) {
	http.SetCookie(w, &http.Cookie{
		Name: "counter",
		Value: strconv.Itoa(counter),
	})
	fmt.Fprintln(w, "you have been to this site",
		counter, "times.")
}
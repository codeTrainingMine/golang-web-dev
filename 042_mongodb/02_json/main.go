package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"golang-web-dev/042_mongodb/02_json/models"
	"encoding/json"
	"fmt"
)

func main()  {
	r := httprouter.New()
	r.GET("/", index)
	// added route plus parameter
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/94873834">GO TO: http://localhost:8080/user/94873834</a>
	</body>
	</html>
	`

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

// changed func name
func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name: "James Bond",
		Gender: "male",
		Age: 32,
		Id: p.ByName("id"),
	}

	// Marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

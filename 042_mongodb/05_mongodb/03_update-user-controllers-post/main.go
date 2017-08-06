package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
	"golang-web-dev/042_mongodb/05_mongodb/03_update-user-controllers-post/controllers"
)

func main()  {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://trash")

	// Check if connection erro, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
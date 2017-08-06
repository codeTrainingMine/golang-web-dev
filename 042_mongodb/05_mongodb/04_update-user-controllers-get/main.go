package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
	"golang-web-dev/042_mongodb/05_mongodb/04_update-user-controllers-get/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://trash")

	if err != nil {
		panic(err)
	}
	return s
}
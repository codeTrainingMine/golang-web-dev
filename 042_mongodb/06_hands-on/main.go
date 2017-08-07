package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"golang-web-dev/042_mongodb/06_hands-on/controllers"
	"golang-web-dev/042_mongodb/06_hands-on/models"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getMap())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getMap() map[bson.ObjectId]models.User {
	myMap := map[bson.ObjectId]models.User{}
	return myMap
}

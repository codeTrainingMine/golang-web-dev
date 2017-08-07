package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"golang-web-dev/042_mongodb/08_hands-on/controllers"
	"golang-web-dev/042_mongodb/08_hands-on/models"
	"io/ioutil"
	"fmt"
	"encoding/json"
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

func getMap() map[string]models.User {
	myMap := map[string]models.User{}

	dat, err := ioutil.ReadFile("saved-map")
	check(err)

	fmt.Println("1")
	if err == nil {
		fmt.Println("0")
		err = json.Unmarshal(dat, &myMap)
		check(err)

		for k, v := range myMap {
			fmt.Println(k, v)
		}
	}

	return myMap
}

func check(e error)  {
	if e != nil {
		fmt.Println(e)
	}
}

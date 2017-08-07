package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"golang-web-dev/042_mongodb/08_hands-on/models"
	"os"
	"github.com/satori/go.uuid"
)

type UserController struct {
	myMap map[string]models.User
}

func NewUserController(s map[string]models.User) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// composite literal
	u := models.User{}

	// Fetch user
	u, ok := uc.myMap[id]
	if !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = uuid.NewV4().String()

	uc.myMap[u.Id] = u

	uc.saveFile()

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) saveFile()  {
	mj, _ := json.Marshal(uc.myMap)

	f, err := os.Create("saved-map")
	check(err)

	defer f.Close()

	_, err = f.Write([]byte(mj))
	check(err)
}

func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Delete user
	u, ok := uc.myMap[id]
	if !ok {
		w.WriteHeader(404)
		return
	}

	delete(uc.myMap, u.Id)

	uc.saveFile()

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}

package main

import (
	//"github.com/satori/go.uuid"
	//"golang.org/x/crypto/bcrypt"
	"net/http"
	"golang-web-dev/042_mongodb/10_hands-on/controllers"
)

func main() {
	uc := controllers.NewUserController()
	http.HandleFunc("/", uc.Index)
	http.HandleFunc("/bar", uc.Bar)
	http.HandleFunc("/signup", uc.Signup)
	http.HandleFunc("/login", uc.Login)
	http.HandleFunc("/logout", uc.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}


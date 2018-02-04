package config

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

// database
var DB *mgo.Database

// collections
var Books *mgo.Collection

func init()  {
	// get a mongo session
	// s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")
	// mongodb://bond:moneypenny007@localhost:27017/bookstore
	fmt.Println("attempting connect")
	s, err := mgo.Dial("mongodb://test3:mypassbcde3@192.168.1.3:27017/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("You connected to your mongo database.")
}
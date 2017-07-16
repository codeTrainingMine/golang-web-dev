package main

import (
	"database/sql"
	"net/http"
	"io"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main()  {
	db, err = sql.Open("mysql", "awsuser:mypassword(mydbinstance.ct4eaa39mh3r.ap-southeast-2.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request)  {
	_, err = io.WriteString(w, "Successfully complete.")
	check(err)
}

func check(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}
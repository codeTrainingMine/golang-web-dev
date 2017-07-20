package main

import (
	"database/sql"
	"net/http"
	"io"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB
var err error
var hostname string

func init()  {
	hostname, _ = os.Hostname()
}
func main()  {
	db, err = sql.Open("mysql", "awsuser:mypassword@tcp(mydbinstance.ct4eaa39mh3r.ap-southeast-2.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/ping", ping)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":80", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintln(w, hostname, "at index")
}

func amigos(w http.ResponseWriter, req *http.Request)  {
	rows, err := db.Query(`select aName from amigos;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, hostname, s)
}

func create(w http.ResponseWriter, req *http.Request)  {
	stmt, err := db.Prepare(`create table customer (name varchar(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, hostname, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request)  {
	stmt, err := db.Prepare(`insert into customer values ("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)
	
	n, err := r.RowsAffected()
	check(err)
	
	fmt.Fprintln(w, hostname, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request)  {
	rows, err := db.Query(`select * from customer;`)
	check(err)
	defer rows.Close()
	
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, hostname, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request)  {
	stmt, err := db.Prepare(`update customer set name = "Jimmy" where name = "James";`)
	check(err)
	defer stmt.Close()
	
	r, err := stmt.Exec()
	check(err)
	
	n, err := r.RowsAffected()
	check(err)
	
	fmt.Fprintln(w, hostname, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request)  {
	stmt, err := db.Prepare(`delete from customer where name="Jimmy";`)
	check(err)
	defer stmt.Close()
	
	r, err := stmt.Exec()
	check(err)
	
	n, err := r.RowsAffected()
	check(err)
	
	fmt.Fprintln(w, hostname, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request)  {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, hostname, "DROPPED TABLE customer")
}

func check(err error)  {
	if err!= nil {
		fmt.Println(err)
	}
}

func ping(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "OK")
}
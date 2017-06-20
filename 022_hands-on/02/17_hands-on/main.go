package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"bufio"
	"strings"
)

func main()  {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		co, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleCon(co)
	}
}

func handleCon(co net.Conn)  {
	defer co.Close()

	sc := bufio.NewScanner(co)

	counter := 0
	var reqLine string
	for ; sc.Scan() ; {
		ln := sc.Text()
		if counter == 0 {
			reqLine = ln
		}
		fmt.Println(ln)
		if ln == "" {
			break
		}
		counter++
	}

	words := strings.Split(reqLine, " ")
	method := words[0]
	path := words[1]
	fmt.Println(method)
	fmt.Println(path)

	var data string
	switch {
	case method == "GET" && path == "/":
		data = "you requested a GET /"
	case method == "GET" && path == "/apply":
		data = "you requested a GET /apply"
	case method == "POST" && path == "/apply":
		data = "you requested a POST /apply"
	default:
		data = "returning the default"
	}

	body :=
		`<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Title</title>
		</head>
		<body>
		<h1>` + data + `</h1>
		<form method="post" action="/apply">
		<input type="submit" value="test">
		</form>
		</body>
		</html>`
	io.WriteString(co, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(co, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(co, "Content-Type: text/html\r\n")
	io.WriteString(co, "\r\n")
	io.WriteString(co, body)
}
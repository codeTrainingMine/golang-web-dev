package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"bufio"
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

	for ; sc.Scan() ; {
		ln := sc.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}

	body :=
		`<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Title</title>
		</head>
		<body>
		<h1>HOLY COW THIS IS LOW LEVEL</h1>
		</body>
		</html>`
	io.WriteString(co, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(co, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(co, "Content-Type: text/html\r\n")
	io.WriteString(co, "\r\n")
	io.WriteString(co, body)
}
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main()  {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn)  {
	defer conn.Close()

	// read request
	method, uri := request(conn)

	switch method {
	case "GET":
		switch uri {
		case "/Apply":
			respondGetApply(conn)
		case "/Submit":
			respondGetSubmit(conn)
		default:
			respond(conn, "You called GET something else")
		}
	case "POST":
		switch uri {
		case "/Apply":
			respondPostApply(conn)
		case "/Submit":
			respondPostSubmit(conn)
		default:
			respond(conn, "You called POST  something else")
		}
	}

	// write response
	//respond(conn, uri)
}

func respondGetApply(conn net.Conn)  {
	respond(conn, "You called GET /Apply")
}

func respondGetSubmit(conn net.Conn)  {
	respond(conn, "You called GET /Submit")
}

func respondPostApply(conn net.Conn)  {
	respond(conn, "You called POST /Apply")
}

func respondPostSubmit(conn net.Conn)  {
	respond(conn, "You called POST /Submit")
}

func request(conn net.Conn) (string, string) {
	i := 0
	var method, uri string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			method = strings.Fields(ln)[0]
			fmt.Println("***METHOD", method)
			uri = strings.Fields(ln)[1]
			fmt.Println("***URI", uri)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return method, uri
}

func respond(conn net.Conn, body string)  {

	b := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>` + body + `</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(b))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, b)
}
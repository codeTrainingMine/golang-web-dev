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

	body := "<h1>hello world</h1><br>" + reqLine
	io.WriteString(co, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(co, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(co, "Content-Type: text/plain\r\n")
	io.WriteString(co, "\r\n")
	io.WriteString(co, body)
}
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
	}

	io.WriteString(co, "I see you connected.\n")

}
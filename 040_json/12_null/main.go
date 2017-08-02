package main

import (
	"encoding/json"
	"log"
	"fmt"
)

func main()  {
	//var a []string
	//var a string
	var a int

	rcvd := `null`
	err := json.Unmarshal([]byte(rcvd), &a)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a)
	//fmt.Println(len(a))
	//fmt.Println(cap(a))
}
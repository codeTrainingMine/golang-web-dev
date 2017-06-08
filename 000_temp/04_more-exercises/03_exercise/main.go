package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main()  {
	p1 := person{"Joe", "Smith"}
	fmt.Println(p1)
	fmt.Println(p1.fname)
}
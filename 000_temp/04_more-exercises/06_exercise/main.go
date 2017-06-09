package main

import "fmt"

type person struct {
	fname string
	lname string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fname, "is walking.")
}

func main()  {
	p1 := person{
		"Joe",
		"Smith",
		[]string{"Burger", "Pizza", "Rice"},
	}

	s := p1.walk()
	fmt.Println(s)
}
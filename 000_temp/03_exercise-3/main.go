package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak()  {
	fmt.Println("I'm a person and my name is ", p.fname)
}

func (s secretAgent) speak()  {
	fmt.Println("I'm a secret agent and my name is ", s.fname)
}

type human interface {
	speak()
}

func saySomething(h human)  {
	h.speak()
}

func main()  {
	p := person{
		"Joe",
		"Smith",
	}

	s := secretAgent{
		person {
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p)
	saySomething(s)
}
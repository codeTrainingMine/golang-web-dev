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

func (p person) pSpeak()  {
	fmt.Println("I'm a person and my name is ", p.fname)
}

func (s secretAgent) saSpeak()  {
	fmt.Println("I'm a secret agent and my name is ", s.fname)
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
	fmt.Println(p.fname)
	p.pSpeak()
	fmt.Println(s.fname)
	s.saSpeak()
	s.person.pSpeak()
}
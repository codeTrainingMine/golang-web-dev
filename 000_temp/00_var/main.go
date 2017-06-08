package main

import "fmt"

var y int
var z int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak()  {
	fmt.Println(p.fname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak()  {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human)  {
	h.speak()
}

func main()  {
	x := 7
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y = 8
	fmt.Println(y)

	fmt.Println(z)

	xi := []int{2, 4, 6, 9, 42}
	fmt.Println(xi)

	m := map[string]int {
		"Todd": 45,
		"Job": 42,
	}
	fmt.Println(m)

	p := person{"Joe", "Smith"}
	fmt.Println(p, p.lname)

	p.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa1.person.speak()

	saySomething(sa1)
	saySomething(p)
}
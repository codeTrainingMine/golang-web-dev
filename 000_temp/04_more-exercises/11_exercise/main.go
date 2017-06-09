package main

import "fmt"

type gator int

var g1 gator = 5

func (g gator) greeting()  {
	fmt.Println("Hello, I am a gator")
}

type flamingo bool

func (f flamingo) greeting()  {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

type swampCreature interface {
	greeting()
}

func bayou(s swampCreature)  {
	s.greeting()
}

func main()  {
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//x = g1 // error different types
	x = int(g1)
	fmt.Println(x)

	var f1 flamingo = true

	//g1.greeting()
	bayou(g1)
	bayou(f1)
}
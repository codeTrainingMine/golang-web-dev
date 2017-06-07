package main

import "fmt"

var y int
var z int

type person struct {
	fname string
	lname string
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
}
package main

import "fmt"

func main()  {
	m := map[string]int{"John":10, "Mark":12, "Peter":7}
	fmt.Println(m)

	for n := range m {
		fmt.Println(n)
	}

	for n, v := range m {
		fmt.Println(n, v)
	}
}
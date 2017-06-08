package main

import "fmt"

func main()  {
	xi := []int{1, 2, 3, 4, 7, 11}
	fmt.Println(xi)

	for n, _ := range xi {
		fmt.Println(n)
	}

	for n, v := range xi {
		fmt.Println(n, v)
	}
}
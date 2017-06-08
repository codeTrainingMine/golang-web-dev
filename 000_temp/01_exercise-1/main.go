package main

import (
	"math"
	"fmt"
)

type square struct {
	side float64
}

func (s square) area() float64  {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64  {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape)  {
	fmt.Println(s.area())
}

func main()  {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}
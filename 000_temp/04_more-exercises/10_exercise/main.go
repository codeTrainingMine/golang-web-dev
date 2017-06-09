package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

func (t truck) transportationDevice() string  {
	return "transports cargo"
}

type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() string  {
	return "transports people"
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation)  {
	fmt.Println(t.transportationDevice())
}

func main()  {
	t := truck{
		vehicle{
			4,
			"white",
		},
		true,
	}
	v := sedan{
		vehicle{
			4,
			"red",
		},
		false,
	}

	report(t)
	report(v)
}
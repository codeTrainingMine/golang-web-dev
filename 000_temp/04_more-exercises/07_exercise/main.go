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
	fmt.Println(t)
	fmt.Println(t.color)
	fmt.Println(v)
	fmt.Println(v.luxury)
	fmt.Println(t.transportationDevice())
	fmt.Println(v.transportationDevice())
}
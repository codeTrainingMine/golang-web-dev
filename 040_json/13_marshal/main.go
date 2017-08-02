package main

import (
	"fmt"
	"encoding/json"
)

type model struct {
	State bool
	Pictures []string
}

func main()  {
	m := model{}

	fmt.Println(m)

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(string(bs))
}
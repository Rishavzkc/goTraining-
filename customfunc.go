package main

import (
	"fmt"
)

//custom type and receiver function

type cities []string

func main() {

	India := cities{"Pune", "Mumbai", "Delhi", "Kolkata"}
	India.print()
}

func (c cities) print() {
	for i, city := range c {
		fmt.Println(i, city)
	}
}

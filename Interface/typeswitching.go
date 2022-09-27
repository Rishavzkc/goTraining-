package main

import (
	"fmt"
	"strings"
)

func explain(i interface{}) {
	switch i.(type) { //here we use type switching for checking type of interface
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string))) //type assertion i.e. assert any datatype in interface
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}

func main() {
	explain("Hello")
	explain(50)
	explain(true)
}

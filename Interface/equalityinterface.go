package main

import "fmt"

func Equality(i interface{}, j interface{}) {
	if i == j {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
}

func main() {
	var i interface{}
	var j interface{}

	Equality(i, j)

	var a interface{} = "Hi"
	var b interface{} = "Hello"

	Equality(a, b)
}

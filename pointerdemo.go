package main

import "fmt"

func main() {

	var a *int //integer pointer

	b := 5
	a = &b               // reference b memory location
	fmt.Printf("%v", *a) //a is holding b value

}

package main

import "fmt"

func main() {
	x := 7
	y := &x

	fmt.Println(x, y)
	*y = 9
	fmt.Println(x, y) //here x and y pointing same memory location

	tochange := "hello"
	var point *string = &tochange
	fmt.Println(point)
	fmt.Println(&point)

}

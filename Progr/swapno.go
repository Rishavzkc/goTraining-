package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 20

	fmt.Println("No. before swaping:x=  ", x)
	fmt.Println("No. before swapping:y= ", y)

	x = x + y
	y = x - y
	x = x - y

	fmt.Println("No. after swaping: x= ", x)
	fmt.Println("No. after swapping: y=", y)

}

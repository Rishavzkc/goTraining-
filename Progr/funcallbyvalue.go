package main

import "fmt"

func main() {

	var a int = 100
	var b int = 200

	fmt.Println("No before swaping a\n ", a)
	fmt.Println("No before swaping b\n ", b)

	swap(a, b)
}

func swap(x int, y int) int {

	var temp int

	temp = x
	x = y
	y = temp

	fmt.Println("No before swaping a\n ", x)
	fmt.Println("No before swaping b\n ", y)

	return temp

}

package main

import "fmt"

func main() {
	summation(2, 3, 4, 5, 6, 7, 8)

	sum, sub := calc(19, 4)
	fmt.Println(sum, sub)
}

//here we use 3 dots to take any number of variable
func summation(vals ...int) {
	sum := 0
	for _, n := range vals {
		sum += n
	}
	fmt.Println(sum)
}

func calc(x int, y int) (int, int) {
	return (x + y), (x - y)
}

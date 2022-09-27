package main

import "fmt"

func main() {

	num1 := []int{1, 2, 3, 4, 5}
	num2 := []int{2, 4, 5, 6, 7}

	for i := 0; i < len(num1); i++ {
		sum := num1[i] + num2[i]

		fmt.Println(sum)

	}
}

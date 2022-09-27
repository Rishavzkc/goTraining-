package main

import (
	"fmt"
)

func main() {

	num1 := []int{1, 2, 3, 4, 5, 6}
	//printing num1 slice
	fmt.Println(num1)

	//reverse of num1
	num3 := slicefunc(num1)
	fmt.Println(num3)

	//Adding two slice

	sum := [7]int{}

	for j := 0; j < len(num1); j++ {
		sum[j] = num1[j] + num3[j]
		fmt.Println("Sum : ", sum[j])
	}

}

func slicefunc(num1 []int) []int {
	var num2 []int

	for i := len(num1) - 1; i >= 0; i-- {

		num2 = append(num2, num1[i])
	}
	return num2

}

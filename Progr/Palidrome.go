package main

import "fmt"

func main() {
	num := 121
	var r int
	sum := 0

	original := num

	for num > 0 {
		r = num % 10
		num = num / 10
		sum = (sum * 10) + r

	}
	fmt.Println(sum)
	if original == sum {
		fmt.Println("Number is palidrome")
	} else {
		fmt.Println("Not palidrome")
	}

}

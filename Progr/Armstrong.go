package main

import "fmt"

func main() {

	//153 = 1^3+ 5^3+ 3^3

	num := 153
	temp := 0
	sum := 0
	original := num
	for num > 0 {
		temp = num % 10
		num = num / 10
		sum = sum + temp*temp*temp
	}

	if original == sum {
		fmt.Println("Armstrong Number")
	} else {
		fmt.Println("Not Armstrong Number")
	}
}

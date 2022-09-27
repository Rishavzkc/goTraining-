package main

import "fmt"

func main() {
	a := [5]int{20, 45, 67, 89, 32}

	var temp int

	size := len(a)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if a[i] > a[j] {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println("Second largest Number: ", a[size-2])
}

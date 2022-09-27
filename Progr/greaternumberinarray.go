package main

import "fmt"

func main() {
	a := [6]int{13, 45, 67, 88, 98, 56}

	max := a[0]

	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]

		}
	}
	fmt.Println(max)
}

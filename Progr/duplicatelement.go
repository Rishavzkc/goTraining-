package main

import (
	"fmt"
)

func main() {
	var a []int = []int{1, 2, 4, 5, 67, 7, 2, 4}

	/*	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				fmt.Println(a[j])
			}
		}

	} */

	for k, element := range a {
		for l, element2 := range a {
			if element == element2 && k > l {
				fmt.Println(element)
			}
		}
	}
}

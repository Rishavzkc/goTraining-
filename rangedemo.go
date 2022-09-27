package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 4, 5, 6, 8}

	//iterating using for loop

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	//iterating using range
	for i, element := range a {
		//fmt.Println(i,element)
		fmt.Printf("%d %d\n", i, element)

	}

	for _, value := range a {
		fmt.Printf("%d\n", value)
	}
}

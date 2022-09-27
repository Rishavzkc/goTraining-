package main

import (
	"fmt"
)

func main() {

	//declaring and initializing array
	//M1-
	var arr [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	//M2
	var arr2 = [...]int{1, 2, 3, 45, 5, 8}
	fmt.Println(arr2)

	//M3
	arr3 := [...]int{1, 2, 34, 5}
	fmt.Println(arr3)

	//length of array
	arr4 := arr2
	fmt.Println(arr4)
	fmt.Println(len((arr4)))

	//multidimentional array
	var matrix [3][3]int = [3][3]int{{1, 2, 3}, {2, 3, 4}, {5, 6, 7}}
	fmt.Println(matrix)
}

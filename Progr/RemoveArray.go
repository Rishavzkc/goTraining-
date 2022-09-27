package main

import (
	"fmt"
)

func main() {

	// initialize an array
	array := [5]int{11, 22, 33, 44, 55}
	fmt.Println("The original array is:", array)

	// initialize the index of the element to delete
	i := 2

	// check if the index is within array bounds

	if i < 0 || i >= len(array) {
		fmt.Println("The given index is out of bounds.")
	} else {

		// delete an element from the array
		array2 := 0

		for index := range array {
			if i != index {
				array[array2] = array[index]
				array2++
			}
		}

		// reslice the array to remove extra index
		newArray := array[:array2]
		fmt.Println("The new array is:", newArray)
	}
}

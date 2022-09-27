package main

import "fmt"

func main() {
	var arr [5]int

	arr[0] = 100
	arr[4] = 50

	arr[2] = 56
	fmt.Println(arr)

	//sum of array
	array := [3]int{1, 4, 6}
	sum := 0

	for i := 0; i < len(array); i++ {
		sum += array[i]

	}
	fmt.Println(sum)
}

package main

import "fmt"

// input -{1,0,2,0,3,0,5,0,8,0,9}
//output-{1,2,3,4,5,8,0,0,0,0,0,}

func main() {
	var arr []int = []int{1, 0, 2, 0, 3, 0, 5, 0, 8, 0, 9}

	length := len(arr)
	count := 0
	for i := 0; i < length; i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
		}
	}
	for count < length {
		arr[count] = 0
	}
	for k := 0; k < length; k++ {
		fmt.Println(arr[k], " ")
	}

}

package main

import "fmt"

func main() {

	var num1 [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	var num2 [][]int = [][]int{{9, 8, 7}, {6, 5, 4}}

	//var sum [][][]int

	fmt.Println(num1)
	fmt.Println(num2)

	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num1); j++ {

			res := num1[i][j] + num2[i][j]
			fmt.Print(res)

			fmt.Println()
		}
	}
}

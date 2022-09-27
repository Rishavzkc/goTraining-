package main

import "fmt"

func main() {

	ans := 3

	switch ans {
	case 1:
		fmt.Println("Case 1 Passed ")

	case 2:
		fmt.Println("Case 2 Passed ")

	case 5:
		fmt.Println("Case 5 Passed ")

	default:
		fmt.Println("No case passed ")
	}

}

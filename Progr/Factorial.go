package main

import (
	"fmt"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("Enter the Number: \n")

	//	scanner.Scan()
	//	num := scanner.Text()
	num := 5
	result := 1

	for num > 0 {
		result = result * num
		num--
	}

	fmt.Println("Factorial is : ", result)
}

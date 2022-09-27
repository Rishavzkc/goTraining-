package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the Number: \n")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Enter the Name: \n")
	scanner.Scan()
	name:=scanner.Text()


	//	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) //converting string into int

	fmt.Println("Your Result is : ", input)
	fmt.Println("Your Name is : ", name)
}

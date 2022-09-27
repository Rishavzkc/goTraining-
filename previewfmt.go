package main

import (
	"fmt"
)

func main() {

	//Printf print the formatted string
	fmt.Printf("Hello %T %v ", 10, 20)

	//Sprintf store the value in the variable
	var x string = fmt.Sprintf("Hi %T %v", 50, 60)
	fmt.Printf(x)
	fmt.Println()

	//Number in binary
	fmt.Printf("Number : %b", 456)

	fmt.Println()
	//Displaying floating number with scientific notation
	fmt.Printf("Number : %e", 8976.9)
	fmt.Println()

	//formating string
	fmt.Printf("Number : %s", "Word")
	fmt.Println()

	//formating string and print in double quoted
	fmt.Printf("Number : %q", "Word")
	fmt.Println()

	//Width and percision  And will print 9 space before trim from leftside
	fmt.Printf("Number: %9.2q", "hello")
	fmt.Println()
	//padding (here 45 will print after 7 zero -o/p --000000045)
	fmt.Printf("Number: %09d", 45)
}

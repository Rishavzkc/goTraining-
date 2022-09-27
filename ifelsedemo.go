package main

import (
	"fmt"
)

func main() {
	age := 19
	fmt.Println("Condition for driving")
	if age >= 18 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("Wait untill year", 18-age)
	}

}

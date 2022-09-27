package main

import (
	"fmt"
)

func main() {

	//for only with condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//For with initialization/condition/increment
	for j := 7; j <= 12; j++ {
		fmt.Println(j)
	}

	//Infinite for loop with break
	for {
		fmt.Println("loop")
		break
	}

	//For with continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Continue", n)
	}
}

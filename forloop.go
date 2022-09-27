package main

import "fmt"

func main() {
	//1st way
	/*	x := 3

		for x < 8 {
			fmt.Println(x)
			x++
		}
		// 2nd way
		for i := 0; i < 8; i++ {
			fmt.Println(i)
		} */

	//Number divisible by 3,7 and 9
	for i := 0; i < 1000; i++ {
		if i != 0 && i%3 == 0 && i%7 == 0 && i%9 == 0 {
			fmt.Println(i)
			continue
		}
	}

}

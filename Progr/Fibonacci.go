package main

import "fmt"

func main() {

	//1 1 2 3 5 8 13 21

	num := 10

	f1 := 0
	f2 := 0
	f3 := 1

	for i := 1; i <= num; i++ {
		fmt.Println(f3)
		f1 = f2
		f2 = f3
		f3 = f1 + f2
	}

}

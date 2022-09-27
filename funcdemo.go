package main

import "fmt"

func main() {
	test := func(x int) int {
		return x * -1
	}
	fmt.Println(test(8))

	test2(test)
}

func test2(myfunc func(int) int) {
	fmt.Println(myfunc(7))

}

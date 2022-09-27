package main

import "fmt"

func changefunc(slice []int) {
	slice[1] = 500
}

func main() {
	var x []int = []int{12, 34, 56, 78}

	fmt.Println(x)
	changefunc(x) //when we passing value of x in function means slice =x & x=pointer to above slice
	fmt.Println(x)
}

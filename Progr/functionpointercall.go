package main

import "fmt"

func DoSomething(a *int) {

	*a = 78
}

func main() {

	x := 89
	fmt.Println(x)
	DoSomething(&x)

	fmt.Println(x)
}

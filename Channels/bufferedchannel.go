package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 5

	fmt.Println("Sending value")

	val := <-ch
	fmt.Println("Receiving value = ", val)
}

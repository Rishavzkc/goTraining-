package main

import (
	"fmt"
)

func send(ch chan int) {
	ch <- 3
}

func receive(ch chan int) {

	val := <-ch

	fmt.Printf("Received value = %d ", val)
}

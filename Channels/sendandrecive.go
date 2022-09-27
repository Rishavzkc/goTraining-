package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	fmt.Println("Sending value to channel ")
	go send(ch)

	fmt.Println("Reciving value from channel")
	go receive(ch)

	time.Sleep(1 * time.Second)

}

func send(ch chan int) {
	ch <- 3
}

func receive(ch chan int) {

	val := <-ch

	fmt.Printf("Received value = %d ", val)
}

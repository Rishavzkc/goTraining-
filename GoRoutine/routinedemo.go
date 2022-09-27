package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Before calling go routine")

	go start()
	time.Sleep(1 * time.Second)
	fmt.Println("After calling go routine")
}
func start() {
	fmt.Println("In goroutine")
}

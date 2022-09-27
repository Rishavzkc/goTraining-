package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Statement 1")
	go start()
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}

func start() {
	fmt.Println("Statement 2")
	go start2()
}

func start2() {
	fmt.Println("Statement 3")
}

//the first goroutine starts the second goroutine.
// The first goroutine then prints “Statement 2” and then it exits.
// The second goroutine then starts and prints “statement 3”
//goroutines don’t have parents or children and they exist as an independent execution.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	for i := 0; i < 10; i++ {

		go execute(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Finish")
}

func execute(id int) {
	fmt.Printf("id: %d \n", id)
}

//Every time when we run program it will give different output.
// since the goroutines will be run concurrently and it is not determined which will run first.

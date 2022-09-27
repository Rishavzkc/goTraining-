package main

import (
	"fmt"
	"time"
)

func AddNumber(number int, ch chan int) {
	result := number + 10
	ch <- result

}

func main() {
	fmt.Println("Executing GOroutine")

	ch := make(chan int)
	go AddNumber(7, ch)
	fmt.Println("Done")
	time.Sleep(1 * time.Second)

	result := <-ch

	fmt.Printf("Result value is %v, and type is %T ", result, result)
}

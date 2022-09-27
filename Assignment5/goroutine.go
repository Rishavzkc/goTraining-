package main

import (
	"fmt"
	"sync"
)

func AddNumber(wg *sync.WaitGroup, ch chan int, number int) {
	result := number + 10
	ch <- result
	wg.Done()
}

func main() {
	fmt.Println("Executing Goroutine")
	ch := make(chan int, 4)

	var wg sync.WaitGroup
	wg.Add(4)

	go AddNumber(&wg, ch, 7)
	go AddNumber(&wg, ch, 15)
	go AddNumber(&wg, ch, 9)
	go AddNumber(&wg, ch, 4)
	fmt.Println("Done")

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Printf("Result value is %v , and type is %T \n ", result, result)
	}
	fmt.Println()
	if len(ch) == 0 {
		fmt.Println("Empty")
	} else {
		fmt.Println("Not Empty")
	}

}

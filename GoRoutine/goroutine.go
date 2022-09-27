package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// sync package is used for wait group
func main() {

	//to know no. of logical processor
	fmt.Println("Logical Processor: ", runtime.NumCPU())

	//we use Add to say that there is only one go routine
	wg.Add(2)

	//Invoking Anoynomus Function
	go func() {
		fmt.Println("Anonynous Function call")
		wg.Done()
	}()

	go Helloworld()
	fmt.Println("Go routine")

	//wait function will wait for all goroutine to complete execution
	wg.Wait()
}

func Helloworld() {

	//After completing all execution it will tell that all execution is completed
	//we use defer to execute the statement in the last
	defer wg.Done()
	fmt.Println("Hello World ")

	for i := 0; i < 50; i++ {
		fmt.Println(i)
	}

}

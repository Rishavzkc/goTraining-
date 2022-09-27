package main

import "fmt"

func main() {
	var a chan int

	// by using make, an instance of hchan struct is created
	// and all the fields are initialized to their default values

	a = make(chan int)
	//channel is internally represented by a hchan struct
	fmt.Println(a)
}

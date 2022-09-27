package main

import "fmt"

func main() {

	//Here we are using single data varible to store different data type

	var data interface{}

	data = "Hello"
	data = 89
	data = true

	fmt.Println(data)
}

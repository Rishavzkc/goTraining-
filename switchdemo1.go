package main

import "fmt"

func main() {

	var x interface{} = 3

	switch x.(type) {

	case int:
		fmt.Println("Hey")

	case string:
		fmt.Println("Hello")

	case float32:
		fmt.Println("This is float")

	default:
		fmt.Println("This is not valid")
	}

}

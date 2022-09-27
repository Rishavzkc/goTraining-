package main

import "fmt"

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("It is integer")
	case string:
		fmt.Println("This is String")
	default:
		fmt.Println("Others")
	}
}

func main() {
	var i interface{} = "Hi "
	checkType(i)
}

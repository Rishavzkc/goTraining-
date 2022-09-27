package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int32
	a = 45

	var b int64 = int64(a)

	fmt.Println(b)

	fmt.Printf("%v ,%T", b, b)

	//int to string
	var c string = strconv.Itoa(int(a))
	fmt.Println("\n String conversion  ", c)
	fmt.Printf("%v ,%T", c, c)

	//string to int
	val := "567"
	d, _ := strconv.Atoi(val)

	fmt.Println("\n ", d)

	fmt.Printf("%v, %T", d, d)
}

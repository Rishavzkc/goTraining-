package main

import (
	"fmt"
)

func main() {

	var num1 int = 9
	var num2 int = 4

	ans := num1 / num2

	fmt.Println(ans) //op-2 because num1 and num2 are int so o/p is also int

	var num3 float32 = 9
	var num4 float32 = 4

	answer := num3 / num4
	fmt.Println(answer) //op-2 because num3 and num4 are floatso o/p is also float

}

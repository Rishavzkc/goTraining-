package main

import (
	"fmt"
	"strings"
)

func main() {
	// I like java programming language

	str := " I like java programming language"

	str1 := strings.Split(str, " ")

	for i := len(str1) - 1; i >= 0; i-- {
		fmt.Println(str1[i])
	}

}

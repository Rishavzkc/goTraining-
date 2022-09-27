package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("Can not divide by 0")
	}
	return a / b, nil
}

func main() {
	result, err := division(20, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result = ", result)
}

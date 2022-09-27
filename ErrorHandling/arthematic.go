package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(divide(23, 0))
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide through zero")
	}

	return a / b, nil
}

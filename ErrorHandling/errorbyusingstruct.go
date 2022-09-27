package main

import (
	"fmt"
)

type ArthematicError struct {
	textError string
}

func (a ArthematicError) Error() string {
	return a.textError

}

func divideError() error {
	return ArthematicError{"Divide by zero "}
}

func getdivide(num1, num2 int) (int, error) {
	if num2 == 0 {
		return -1, divideError()
	}
	return num1 / num2, nil
}

func main() {

	result, err := getdivide(20, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		fmt.Println("Result = ", result)
	}
}

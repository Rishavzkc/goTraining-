package main

import (
	"errors"
	"fmt"
)

func main() {

	dosomthing()

	c, err := multivalue(2, 3)
	fmt.Println(c)
	fmt.Println(err)

	c1, _ := multivalue(1, 2)
	fmt.Println(c1)

	s := "MyName"
	print(&s)

}

//assign pointer to data type
func print(s *string) {
	fmt.Println(s)

}

func dosomthing() {
	fmt.Println("Hello World")
}

//Type1
func multivalue(a int, b int) (int, error) {
	c := a + b
	if c != 3 {
		return 0, errors.New("Error")
	}
	return c, nil
}

//Type2
//we can also declare return type in argument
func multivaluereturn(a, b int) (c int, err error) {
	c = a + b
	if c != 3 {
		err = errors.New("Error")
	}
	return
}

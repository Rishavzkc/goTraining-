package main

import "fmt"

type Employee struct {
	EmpName string
}

func main() {
	var emp *Employee
	fmt.Println(emp)

	emp = new(Employee)
	fmt.Println(emp)
}

package main

import (
	"fmt"
)

func main() {
	createEmployee()
}

// Employee
type Employee struct {
	EmpId      int
	Name       string
	Department string
}

//create Employee function
func createEmployee() {

	//creating instance of Employee by struct variable
	var emp1 Employee
	emp1.EmpId = 1
	emp1.Name = "John"
	emp1.Department = "Accountant"

	//creating an instance by using struct literal
	emp2 := Employee{
		EmpId:      2,
		Name:       "Jack",
		Department: "Software",
	}

	//if order is known field indentifier can be omitted
	emp3 := Employee{
		2,
		"Jack",
		"Software",
	}

	//can initailize few fields
	emp4 := Employee{
		Name: "Ram",
	}

	fmt.Println(fmt.Sprintf("Employee1 is: %+v\n ", emp1))
	fmt.Println(fmt.Sprintf("Employee2 is: %+v\n ", emp2))
	fmt.Println(fmt.Sprintf("Employee3 is: %+v\n ", emp3))
	fmt.Println(fmt.Sprintf("Employee4 is: %+v\n ", emp4))

	if emp1 == emp2 {
		fmt.Println("Both are equal")
	} else {
		fmt.Println("Both are not equal")
	}

	if emp2 == emp3 {
		fmt.Println("Both are equal")
	} else {
		fmt.Println("Both are not equal")
	}
}

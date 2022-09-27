package main

import (
	"fmt"
)

func main() {

	//value receiver
	emp := Employee{
		Name:   "Mohan",
		EmplId: 3456,
	}
	emp.EmployeeDetails()

	//change name
	emp.ChangeName("Ram")
	emp.EmployeeDetailsAfterChange()
}

type Employee struct {
	Name   string
	EmplId int
}

//Employee Details
func (emp Employee) EmployeeDetails() {
	fmt.Println("Employee name is: ", emp.Name)
	fmt.Println("Employee id is: ", emp.EmplId)
}

//Employeee Details after change
func (emp Employee) EmployeeDetailsAfterChange() {
	fmt.Println("Employee name is: ", emp.Name)
	fmt.Println("Employee id is: ", emp.EmplId)
}

//Change name of Employee
func (emp Employee) ChangeName(name string) {
	emp.Name = name
}

package main

import (
	"fmt"
)

func main() {
	//pointer Reciver
	ptrEmp := &Employee{
		Name:   "Jack",
		EmplId: 9879,
	}
	ptrEmp.EmployeeDetails()

	//change name
	ptrEmp.ChangeName("Jonny")
	ptrEmp.EmployeeDetailsAfterChange()
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

//using pointer reciver
//Change name of Employee
func (emp *Employee) ChangeName(name string) {
	emp.Name = name
}

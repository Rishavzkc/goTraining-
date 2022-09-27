package main

import "fmt"

func main() {
	createEmployee()
}

//creating address struct
type Address struct {
	Street string
	City   string
}

//creating employee struct
type Employee struct {
	Empld   int
	Name    string
	Address Address
}

//createEmployee function
func createEmployee() {

	emp := Employee{
		Empld: 23,
		Name:  "Mohan",
		Address: Address{
			Street: "Baner",
			City:   "Pune",
		},
	}
	fmt.Println(fmt.Sprintf("Embeddeed Stuct: %+v\n", emp))

	//printing individual value
	fmt.Println("Name is: ", emp.Name)
}

package main

import (
	"fmt"
)

func main() {
	createStudent()
}

func createStudent() {

	//initializing map using make function
	studentMap := make(map[int]string)

	studentMap[1] = "Jack"
	studentMap[2] = "John"

	//fmt.Println(studentMap)

	for key, value := range studentMap {
		fmt.Println("Student id: ", key, "and Student Name: ", value)
	}

	//initilizing map using map literal
	students := map[int]string{
		1: "Mohan",
		2: "Ram",
	}

	for key, value := range students {
		fmt.Println("Student id is: ", key, "and student name is: ", value)
	}
	//deleting student2
	delete(students, 2)

	//prining single value
	fmt.Println("Name of Student 1 is: ", students[1])

	//if statement
	if _, ok := students[2]; ok {
		fmt.Println("\n Student Found")
	} else {
		fmt.Println("\n Student Not Found")
	}
}

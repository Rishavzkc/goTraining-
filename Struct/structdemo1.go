package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func main() {
	/*	var p1 Point = Point{1, 2}

		fmt.Println("Point: ", p1)

		fmt.Println(p1.x) */

	p2 := &Point{y: 7}
	fmt.Println(p2)
	changeX(p2)
	fmt.Println(p2)

	createstudent()
	Details()
}

func changeX(pt *Point) {
	pt.x = 500
	//fmt.Println(pt)
}

type student struct {
	Name   string
	Id     int
	Level  string
	Branch Branch
}
type Branch struct {
	Branchcode int
	BranchName string
	BranchId   int
}

func Details() {
	studr := student{
		Name:  "john",
		Id:    899,
		Level: "Medium",
		Branch: Branch{
			Branchcode: 909,
			BranchName: "Electrical",
			BranchId:   90,
		},
	}
	fmt.Println("Student Details: ", studr)
}

func createstudent() {
	var stu student
	stu.Name = "John"
	stu.Id = 34
	//	stu.Class = "IT"

	fmt.Println("Student is: ", stu)

	stu1 := student{
		Name: "Jack",
		Id:   567,
		//		Class: "software",
	}
	fmt.Println("Student is: ", stu1)

}

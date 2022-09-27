package main

import "fmt"

func main() {

	c1 := Company{"Hcl", 500, "Pune"}
	c2 := Company{"Quest", 8000, "Bangalore"}

	fmt.Println("Creating Company Details.... ")
	Companies := []CompanyService{c1, c2}
	fmt.Println(Companies)

	fmt.Println("Reading Company Details.... ")
	fmt.Println("Company 1 Name: ", c1.CompanyName)

	fmt.Println("Updating Company Details")
	c1.CompanyName = "TCS"
	fmt.Println(c1)

	fmt.Println("Delting Company Details....")
	c1.CompanyName = " "
	fmt.Println(c1)

}

type CompanyService interface {
	CreateDetails()
	ReadDetails() (string, int)
	UpdateDetails()
	DeleteDetails()
}

type Company struct {
	CompanyName     string
	CompanyId       int
	CompanyLocation string
}

func (c Company) CreateDetails() {
	fmt.Println("Company Name: ", c.CompanyName)
	fmt.Println("Company Id: ", c.CompanyId)
	fmt.Println("Company Location: ", c.CompanyLocation)
}

func (c Company) ReadDetails() (string, int) {
	return c.CompanyName, c.CompanyId
}

func (c Company) UpdateDetails() {

}
func (c Company) DeleteDetails() {

}
func Data(co CompanyService) {
	co.CreateDetails()
}

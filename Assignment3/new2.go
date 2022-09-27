package main

import "fmt"

var Companies = make(map[int]Company)

type CompanyService interface {
	CreateData(companies map[int]Company)
	ReadData(companies map[int]Company)
	UpdateData(companies map[int]Company)
	DeleteData(companies map[int]Company)
}

type Company struct {
	CompanyName     string
	CompanyId       int
	CompanyLocation string
}

func main() {

	var Company CompanyService = New()
	var input int
	for input != 9 {
		fmt.Println("Enter the Company Details")
		fmt.Println("1 For Create")
		fmt.Println("2 For Read")
		fmt.Println("3 For Update")
		fmt.Println("4 For Delete")
		fmt.Println("5 For Exit")

		fmt.Scanln(&input)

		switch input {
		case 1:
			Company.CreateData(Companies)
		case 2:
			Company.ReadData(Companies)
		case 3:
			Company.UpdateData(Companies)
		case 4:
			Company.DeleteData(Companies)
		default:
			fmt.Println("Exit")
		}
	}

}

func (C *Company) CreateData(companies map[int]Company) {
	comp := Company{}
	fmt.Println("Enter Company Name")
	fmt.Scanln(&comp.CompanyName)

	fmt.Println("Enter Company Id")
	fmt.Scanln(&comp.CompanyId)

	fmt.Println("Enter Company Location")
	fmt.Scanln(&comp.CompanyLocation)

	companies[comp.CompanyId] = comp
}

func (C *Company) ReadData(Companies map[int]Company) {
	fmt.Println(Companies)
}

func (C *Company) UpdateData(Companies map[int]Company) {
	var CompName string
	var CompId int
	fmt.Println("Enter Company Name for  updating Id")
	fmt.Scanln(&CompName)

	fmt.Println("Enter Updated Id")
	fmt.Scanln(&CompId)

	updatedComp, _ := Companies[CompId]

	updatedComp.CompanyId = CompId
	Companies[CompId] = updatedComp

}

func (C *Company) DeleteData(company map[int]Company) {
	var id int
	fmt.Println("Enter the Comapany id to delete")
	fmt.Scanln(&id)
	delete(company, id)
}

func New() CompanyService {
	return &Company{}
}

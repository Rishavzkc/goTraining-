package main

var store map[string]Company

type Company struct {
	Id, Name, Location string
}

type CompanyService interface {
	create(Company)
	Update(string, Company)
	Delete(string)
	Read([]Company)
}

func Create(c Company) {
	//store
}

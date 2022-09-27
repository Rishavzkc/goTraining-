package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Company struct {
	Name       string `json:"Name"`
	Id         int    `json:"Id'`
	Location   string `json:"	Location"`
	No_of_Empl int    `json:"	No_of_Empl"`
}

var Companies []Company

func getAllCompanies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Companies)
}

func createCompany(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var company Company
	json.Unmarshal(reqBody, &company)
	Companies = append(Companies, company)
	json.NewEncoder(w).Encode(company)

}

/*func updateComapany(w http.ResponseWriter, r *http.Request) {

	r.Header.Set(Company.Id)
	client := &http.Client{}
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var company Company
	json.Unmarshal(reqBody, &company)
	Companies = append(Companies, company)
	json.NewEncoder(w).Encode(company)
}*/

func deleteCompany(w http.ResponseWriter, r *http.Request) {

	json.Marshal(Companies)
	client := &http.Client{}
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(r.Body)
	var company Company
	json.Unmarshal(reqBody, &company)
	json.NewEncoder(w).Encode("Company is deleted")
}

func apiRequest() {
	http.HandleFunc("/companies", getAllCompanies)
	http.HandleFunc("/company", createCompany)
	//http.HandleFunc("/company{id}", updateComapany)
	http.HandleFunc("/company{id}", deleteCompany)
	log.Fatal(http.ListenAndServe(":8090", nil))
	fmt.Println("Listing on 8090")

}
func main() {
	Companies = []Company{
		Company{Name: "Hcl", Id: 567, Location: "Pune", No_of_Empl: 300},
		Company{Name: "TCS", Id: 789, Location: "Delhi", No_of_Empl: 1000},
	}

	apiRequest()
}

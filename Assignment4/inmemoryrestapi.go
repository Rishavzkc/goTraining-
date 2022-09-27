package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	err := json.NewEncoder(w).Encode(Companies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func createCompany(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var company Company
	unmarshalErr := json.Unmarshal(reqBody, &company)
	if unmarshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Companies = append(Companies, company)
	encodeErr := json.NewEncoder(w).Encode(company)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

/*func updateComapany(w http.ResponseWriter, r *http.Request) {

	r.Header.Set(Company.Id)
	client := &http.Client{}
	resp, _ := client.Do(r)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var company Company
	unmarshalErr := json.Unmarshal(reqBody, &company)
	if unmarshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Companies = append(Companies, company)
	encodeErr := json.NewEncoder(w).Encode(company)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
} */

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	_, marshalErr := json.Marshal(Companies)
	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, _ := client.Do(r)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	reqBody, ioErr := ioutil.ReadAll(r.Body)
	if ioErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var company Company
	unmarshalErr := json.Unmarshal(reqBody, &company)
	if unmarshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeErr := json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{Message: "Company is deleted"})
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func apiRequest() {
	http.HandleFunc("/companies", getAllCompanies)
	http.HandleFunc("/company", createCompany)
	//http.HandleFunc("/company{id}", updateComapany)

	http.HandleFunc("/company/{id}", deleteCompany)
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

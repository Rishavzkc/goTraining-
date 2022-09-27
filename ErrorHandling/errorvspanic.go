package main

import (
	"errors"
	"fmt"
)

func checkSql() bool {
	return false
}

func checkDbConnection() error {
	if checkSql() {
		return nil
	} else {
		return errors.New("My Sql is not Running")
	}
}
func main() {
	if err := checkDbConnection(); err != nil {
		//fmt.Println(err.Error())            //if we use error than even though there is error the flow of program will not stop
		panic("First run MySQL then you can do transaction ") //using panic when there is error the execution of program will stop
	} else {
		fmt.Println("My sql is running")
	}
	fmt.Println("You can do transaction")
}

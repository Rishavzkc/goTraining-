package main

import "fmt"

type error interface {
	Error() string
}

type networkproblem struct {
	Message string
	Id      int
}

//implementing struct and interface method in function
func (np networkproblem) Error() string {
	return fmt.Sprintf("Network Problem! Message: %s ,code: %v ", np.Message, np.Id)
}

func handleError(err error) {
	fmt.Println(err.Error())
}

func main() {
	np := networkproblem{
		Message: "Error Message Recived",
		Id:      404,
	}
	handleError(np)
	//fmt.Println(np)

}

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(file.Name())

}

package main

import "fmt"

func main() {

	f := func() string {
		return ("Hi , All")
	}

	fmt.Println(f())
}

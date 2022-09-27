package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{

		"apple":  5,
		"Mango":  6,
		"Orange": 7,
	}
	fmt.Println(mp)

	names := make(map[string]int)

	names["Ram"] = 1
	names["John"] = 3
	names["jack"] = 4

	fmt.Println(names)

	delete(names, "Ram") //delete one element

	for i, value := range names {
		fmt.Println("Names: ", i, "Id: ", value)
	}

	//to check if key exists
	val, ok := names["jack"]
	fmt.Println(val, ok)

}

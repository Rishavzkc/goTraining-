package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := 100

	fmt.Printf("%v ,%T", a, a)

	var arr [4]int = [4]int{2, 3, 45, 5}
	fmt.Println(arr)

	empl := make(map[string]int)
	empl = map[string]int{
		"Ram":  300,
		"John": 800,
		"jack": 344,
	}
	fmt.Println(empl)

	empl["Ram"] = 4500 //update
	fmt.Println(empl)
	val, ok := empl["ankit"] //if key exists or not
	fmt.Println(val, ok)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You Typed: %q", input)

}

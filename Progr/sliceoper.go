package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 35, 6, 7}

	fmt.Println(cap(arr))

	//coverting array to slice
	var sli []int = arr[1:3]
	//capacity of slice
	fmt.Println(cap(sli))

	fmt.Println(sli[:cap(sli)])

	//slice of int(when we define slice then there is no size)
	var a []int = []int{1, 3, 4, 5, 6}
	b := append(a, 304)
	fmt.Println(b)

	//slice creating by make keyword
	d := make([]int, 5)
	fmt.Printf("%T", d)

	e := make([]string, 8)
	e[0] = "john"
	e[1] = "jack"

	fmt.Println(e)
}

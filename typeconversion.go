package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int32 = 100

	//type conversion int32 to int64
	var b int64 = int64(a)
	fmt.Printf("%v,%T", b, b)
	fmt.Println()

	//type conversion int to string
	var c int = 300
	var str string = strconv.Itoa(c)
	fmt.Printf("%v,%T", str, str)

	//to find real and imagenery number from complex
	d := complex(5, 67)
	fmt.Printf("%v, %v", real(d), imag(d))
	fmt.Println()

	//convert string to byte
	s := "Hello! how are you?"
	bs := []byte(s)
	fmt.Printf("%s", bs)

	fmt.Println(x, y)

}

//iota starts from zero
const (
	x = iota
	_ //here  one value of iota is skipped
	y = iota
)

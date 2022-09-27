package main

import "fmt"

/* func main() {
	a := new(int) //declaring pointer
	b := new(int)

	somepointer(a, b)                        //here compiler does not  ecsape
}
func somepointer(aptr *int, bptr *int) {

} */

func main() {
	a := new(int) //declaring pointer
	b := new(int)

	fmt.Println(a, b)
	//here compiler will escapes to heap means now compiler knows that pointer
	//is not specific to this package
	somepointer(a, b)
}
func somepointer(aptr *int, bptr *int) {

}

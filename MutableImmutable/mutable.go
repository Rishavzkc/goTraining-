package main

import "fmt"

func main() {
	var x []int = []int{1, 3, 56, 78}
	fmt.Println(x)
	y := x

	y[0] = 100

	fmt.Println(y) //if we modify y means modify x

	var mp map[string]int = map[string]int{"hello": 3}

	mp1 := mp
	mp1["hi"] = 40
	fmt.Println(mp, mp1) //if we modify mp1 then mp also get modified

	var z [3]int = [3]int{1, 2, 34}
	k := z

	k[2] = 500

	fmt.Println(z, k) //here if we change the value of k then value of z not change
}

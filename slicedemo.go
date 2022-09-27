package main

import (
	"fmt"
	"strings"
)

func main() {

	//initializing slice with length 4
	s := make([]string, 4)
	fmt.Println("s:", s)

	//adding element to slice based on index
	s[0] = "I"
	s[1] = "am"
	s[2] = "from"
	s[3] = "india"

	fmt.Println("set:", s)
	fmt.Println("at 2nd index =", s[2])

	//to print length of slice
	fmt.Println("length = ", len(s))

	//adding new element to existing slice
	s = append(s, "culture")
	s = append(s, "nice", "place")
	fmt.Println("Appended slice ", s)

	//creating new slice and copying the data from old slice to new slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied data ", c)

	//Re slicing old slice fromindex 2 to 5 ,Re slicing will take elements from >=2 to <5
	re := s[2:5]
	fmt.Println("Re slicing data ", re)

	//taking element from >=0 to <5
	le := s[:5]
	fmt.Println("first 5 element ", le)

	//taking element 2<=
	t := s[2:]
	fmt.Println("element after 1 index ", t)

	//declaring and initializing string
	str := []string{"he", "is", " good", "man"}
	fmt.Println("New string slice ", str)

	//By using join we can seprate with comma
	joinedslice := strings.Join(str, ",")
	fmt.Println("Joined slice: ", joinedslice)

	//if we have slice with 10 element & we want to take 1st 3
	take := s[:3]
	fmt.Println("First three :", take)

	//creating new generics that can take any data type in same slice with help of interface
	generic := make([]interface{}, 4)

	generic[0] = 1
	generic[1] = "String"
	generic[2] = "value"

	fmt.Println("generic value : ", generic)

	//Range operator used for iterate and printing with index value
	//it also give <nil> if element is not present
	for i1, v1 := range generic {
		fmt.Println(i1, v1)
	}

}

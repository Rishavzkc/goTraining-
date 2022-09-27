package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "John, jack, jonny"

	names := strings.Split(str, " ")

	fmt.Println(names[2])
}

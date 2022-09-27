package main

import "fmt"

type Shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func main() {

	c1 := circle{3.4}
	r1 := rectangle{34, 13}

	shapes := []Shape{c1, r1}
	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Println(shape.area())
		fmt.Println(getArea(shape))
	}

}

func (r rectangle) area() float64 {
	return r.breadth * r.length

}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.area()

}

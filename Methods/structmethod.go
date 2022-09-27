package main

import "fmt"

type Student struct {
	name   string
	age    int
	grades []int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) averagegarde() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum) / float64(len(s.grades))
}

func (s Student) maxgrade() int {
	max := 0

	for _, v := range s.grades {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {
	s1 := Student{"john", 20, []int{45, 34, 21, 44}}

	fmt.Println(s1.getAge())

	s1.setAge(23)
	fmt.Println(s1)

	Average := s1.averagegarde()

	fmt.Println(Average)

	max1 := s1.maxgrade()
	fmt.Println(max1)
}

package main

import "fmt"

type Animal struct {
	makeNoiseFn func(*Animal) string
	name        string
	legs        int
}

func (a *Animal) makeNoise() string {
	return a.makeNoiseFn(a)
}

func NewDog(name string) *Animal {
	return &Animal{
		makeNoiseFn: func(a *Animal) string {
			return a.name + "bark!"
		},
		legs: 4,
		name: name,
	}
}

func NewDuck(name string) *Animal {
	return &Animal{
		makeNoiseFn: func(a *Animal) string {
			return a.name + "quack!"
		},
		legs: 4,
		name: name,
	}
}

func main() {
	var dog, duck *Animal

	dog = NewDog("puppy: ")
	duck = NewDuck("donald: ")

	fmt.Println(dog.makeNoise())
	fmt.Println(duck.makeNoise())
}

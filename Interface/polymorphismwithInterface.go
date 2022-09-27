package main

import "fmt"

type Animal interface {
	makeNoise() string
}

type Dog struct {
	Name string
	Legs int
}

func (d *Dog) makeNoise() string {
	return d.Name + "bark!"
}

type Bird struct {
	Name string
	Legs int
}

func (b *Bird) makeNoise() string {
	return b.Name + "quark!"
}

func NewDog(name string) Animal {
	return &Dog{
		Name: name,
		Legs: 4,
	}
}

func NewDuck(name string) Animal {
	return &Bird{
		Name: name,
		Legs: 4,
	}
}

func main() {
	var dog, bird Animal

	dog = NewDog("fido")
	bird = NewDuck("donald")

	fmt.Println(dog.makeNoise())
	fmt.Println(bird.makeNoise())
}

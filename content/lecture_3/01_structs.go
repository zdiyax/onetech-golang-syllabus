package main

import (
	"fmt"
)

const (
	zero = 1 << (8 + iota)
	one
	two
	three
	nine
)

func main() {
	d := Dog{}

	d.name = "Mukhtar"
	d.isReal = true
	d.sound = "woof woof"

	fmt.Println(d)

	d.makeSound()

	fmt.Println(zero)
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(nine)
}

type Animal struct {
	name   string
	sound  string
	isReal bool
}

type Dog struct {
	Animal
}

func (a *Animal) makeSound() {
	fmt.Println(a.sound)
	fmt.Println(a.sound)
	fmt.Println(a.sound)
}

func (d *Dog) makeSound() {
	fmt.Println(d.sound + " POLYMORPHISM " + d.sound)
	fmt.Println(d.sound + " POLYMORPHISM " + d.sound)
	fmt.Println(d.sound + " POLYMORPHISM " + d.sound)
}

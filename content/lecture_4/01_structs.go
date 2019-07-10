package main

import "fmt"

var Far int

func main() {
	//long init
	circle := Circle{Radius: 5, Point: Point{
		x: 0,
		y: 0,
	}}
	//
	//circle := &Circle{Radius: 5, Point:Point{
	//	x:3,
	//	y:3,
	//}}

	//short init
	circle2 := Circle{5, Point{0, 0}}

	fmt.Printf("this is value : %v | this is type : %T\n", circle, circle)

	changeCircleCenterToZeroZero(&circle)

	fmt.Printf("this is value : %v | this is type : %T\n", circle, circle)

	circle2 = circle

	fmt.Printf("this is value : %v | this is type : %T\n", circle2, circle2)

	circle2.x = 100

	fmt.Printf("1 | this is value : %v | this is type : %T\n", circle, circle)
	fmt.Printf("2 | this is value : %v | this is type : %T\n", circle2, circle2)
}

type Point struct {
	x int
	y int
}

type Circle struct {
	//global
	Radius int
	//anonymous field
	Point
}

type Square struct {
	P1 Point
	P2 Point
	P3 Point
	P4 Point
}

func changeCircleCenterToZeroZero(c *Circle) {
	c.x = 0
	c.y = 0
}

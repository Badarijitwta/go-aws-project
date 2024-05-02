package main

import "fmt"

const PI = 3.14

type Circle struct {
	Radius float64
}

func (c *Circle) Circumference() {
	var circumference = 2 * PI * c.Radius
	fmt.Println(circumference)
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		Radius: radius,
	}
}

// func (c *Circle) Area() {
// 	var area float64 = (PI * (c.Radius * c.Radius))
// 	return area
// }

func main() {
	myCicle := NewCircle(2.04)
	myCicle.Circumference()
	fmt.Println(myCicle)
}

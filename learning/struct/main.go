package main

import "fmt"

const PI = 3.14

func main() {

	circ := NewCircle(10)
	circ.Area()
	circ.Circumference()

}

type Circle struct {
	radius float64
}

func (c Circle) Area() {
	area := PI * c.radius * c.radius
	fmt.Println(area)
}

func (c Circle) Circumference() {
	circumference := 2 * PI * c.radius
	fmt.Println(circumference)
}
func NewCircle(radius float64) Circle {
	return Circle{radius: radius}
}

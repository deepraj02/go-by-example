package main

import (
	"fmt"
	"math"
)

func main() {

	a := rectange{height: 3, width: 4}
	b := circle{radius: 5}
	c := square{side: 3}



	measure(a)
	measure(b)

	/// passing pointer value
	measure(&c)

}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// / Defining behavior
type geometry interface {
	area() float64
	perim() float64
}

// / These Struct will implemnt the geometry interface through methods
type rectange struct {
	height, width float64
}
type square struct {
	side float64
}
type circle struct {
	radius float64
}

/// Methods for the Structs

/*
RECTANGLE
*/
func (r rectange) area() float64 {
	return float64(r.height * r.width)
}
func (r rectange) perim() float64 {
	return float64(2*r.height + 2*r.width)
}

/*
CIRCLE
*/
func (c circle) area() float64 {
	return math.Pi * float64(c.radius*c.radius)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * float64(c.radius)
}

/*
SQUARE
*/

func (s square) area() float64 {
	return float64(math.Pow(s.side, 2))
}

func (s square) perim() float64 {
	return float64(4 * s.side)
}

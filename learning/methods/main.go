package main

import "fmt"

func main() {
	r := rect{width: 10, height: 5}
	area := r.area()
	perim := r.perim()
	fmt.Println("area:", area)
	fmt.Println("Perim:", perim)

	/*
	Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	*/
	rp:= &r
	fmt.Println("area:", rp.area())
}

type rect struct {
	width, height int
}

func (r *rect) area() float32 {
	return float32(r.width * r.height)
}

func (r *rect) perim() float32 {
	return float32(2*r.width + 2*r.height)
}

package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perimeter() int {
	return 2*r.height + 2*r.width
}

func main() {

	r := rect{height: 10, width: 15}
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("Area: ", rp.area())
	fmt.Println("Perimeter: ", rp.perimeter())
}

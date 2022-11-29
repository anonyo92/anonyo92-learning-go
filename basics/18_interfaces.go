package basics

import (
	"fmt"
	"math"
)

/* Interfaces are named collections of method signatures */
type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length, width float64
}
type circle struct {
	radius float64
}

/* To implement an interface in Go,
 * we just need to implement all the methods in the interface */

// rectangle implementing the geometry interface
func (r rectangle) area() float64 {
	return r.length * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

// circle implementing the geometry interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

/* If a variable has an interface type, then we can call methods
 * that are in the named interface. Here’s a generic measure function
 * taking advantage of this to work on any geometry type. */
func measure(shape geometry) {
	fmt.Println(shape)
	fmt.Println(shape.area())
	fmt.Println(shape.perimeter())
}

func RunInterfaces() {
	r := rectangle{length: 3, width: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

package basics

import "fmt"

/* Methods define behavior of structs */
type rect struct {
	length, breadth int
}

/* This method has a 'receiver' type of *rect */
func (r *rect) area() int {
	return r.length * r.breadth
}

/* This method has a 'receiver' type of rect */
func (r rect) perimeter() int {
	return 2 * (r.length + r.breadth)
}

func RunMethods() {
	r := rect{length: 20, breadth: 10}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perimeter())

	/* Go automatically handles conversion between values and pointers for method calls.
	 * You may want to use a pointer receiver type
	 * 1. to avoid copying on method calls, or
	 * 2. to allow the method to mutate the receiving struct.
	 */
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perimeter())
}

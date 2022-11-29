package basics

import "fmt"

func runFunctions1() {
	res := sampleFunc1(1, 2)
	fmt.Println("1+2 =", res)
	res = sampleFunc2(1, 2, 3, "a string arg")
	fmt.Println("1+2+3 =", res)
}

/* Simple Function */
func sampleFunc1(a int, b int) int {
	return a + b
}

/* Arg-Types can be shortened if multiple args are of same type */
func sampleFunc2(a, b, c int, d string) int {
	// a, b and c are all ints. d is a string.
	fmt.Println("Inside function sampleFunc2() with d = ", d)
	return a + b + c
}

package basics

import "fmt"

/* Array functions */
func runArrayFuncs() {
	var a = [5]int{10, 20, 30, 40, 50}

	/* 1. len(): Returns length of arr */
	fmt.Println("Length of array: ", len(a))
}

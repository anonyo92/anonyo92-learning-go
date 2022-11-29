package basics

import "fmt"

/* This function returns multiple values */
func sampleFunc3(a, b int) (int, int) {
	return a ^ 2, b ^ 2
}

/* This is "Variadic function" - takes in arbitrary number of arguments */
func sampleFunc4(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func runFunctions2() {
	fmt.Println("Calling a function with mutiple return values:")
	a, b := sampleFunc3(5, 7)
	fmt.Println("It returned these together:", a, b)

	fmt.Println("Calling a variadic function f(1,2); returns", sampleFunc4(1, 2))
	fmt.Println("Calling the same variadic function f(1,2,3); returns", sampleFunc4(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println("nums = ", nums)
	// Can also pass a slice to a variadic function from an existing array like "arr..."
	fmt.Println("Calling the same variadic function f(nums...); returns", sampleFunc4(nums...))
}

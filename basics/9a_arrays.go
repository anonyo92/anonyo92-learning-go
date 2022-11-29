package basics

import "fmt"

func RunArrays() {
	// Declare empty array
	var a [5]int // automatically zero-value initialized
	fmt.Println("emp:", a)

	// Get/Set
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Dynamic initialization
	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println("dcl:", b)

	// 2D array
	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c)

	// Array functions
	fmt.Println("len:", len(a))
	fmt.Println("len:", len(c))
	fmt.Println("len:", len(c[0]))
}

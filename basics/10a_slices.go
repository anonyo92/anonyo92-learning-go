package basics

import "fmt"

func runSlices() {
	var s = make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"

	// Get/Set
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// Indexing/Slicing
	l := s[1:3]
	fmt.Println("slc of s[1:3]:", l)

	l = s[:3]
	fmt.Println("slc of s[:3]:", l)

	l = s[2:]
	fmt.Println("slc of s[2:]:", l)

	// Dynamic initialization
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 2D slice
	var twoD = make([][]int, 3) // first-dim size is 3
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

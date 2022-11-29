package basics

import "fmt"

/* Slice functions */
func runSliceFuncs() {
	s := []string{"a", "b", "c"}
	fmt.Println("dcl:", s)

	/* 1. len(s): returns length of slice */
	fmt.Println("Length: ", len(s))

	/* 2. append(s, *t): appends *t items to s and returns new slice */
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	/* 3. copy(a, b): copies elements of slice s into slice c
	 * Both slices must be of same length */
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	/* 4. */
}

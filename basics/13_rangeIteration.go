package basics

import "fmt"

func RunRangeIter() {
	nums := [4]int{56, 28, -18, 35}
	sum := 0
	/* 1. Iterating through arrays */
	for _, num := range nums {
		// If the index is needed, use a variable instead of _
		sum += num
	}
	fmt.Println("Sum of nums = ", sum)

	/* Iterating through maps */
	// 2a. Using both keys and values
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 2b. Using only keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	/* 3. Iterating through a string */
	/* range on strings iterates over Unicode code points.
	 * The first value is the starting byte index of the 'rune'
	 * and the second the 'rune' itself. */
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

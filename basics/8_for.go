package basics

import "fmt"

func runFor() {
	/* for is go's only looping statement */

	// while-like loop using for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for-like loop using for
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// break
	for { // infinite loop without condition
		fmt.Println("Loop without condition")
		break
	}

	// continue - skip the rest of this iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

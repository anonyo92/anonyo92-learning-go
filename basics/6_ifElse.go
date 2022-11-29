package basics

import "fmt"

func runIfElse() {
	n := 7
	// Simple If
	if n > 0 { // braces are manadatory
		fmt.Println("n is positive")
	}

	// If/Else
	if n%2 == 0 {
		fmt.Println("n is even")
	} else {
		fmt.Println("n is odd")
	}

	// If/Else-if
	if num := 9; num < 0 {
		// 'num' is available within scope of all the branches of this if
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num has only 1 digit")
	} else {
		fmt.Println("num has more than 1 digit")
	}
	// fmt.Print(num) // 'num' is not available here any more

	/* There is no ternary operator in Go */
}

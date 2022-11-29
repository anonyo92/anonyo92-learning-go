package basics

import "fmt"

/* This function returns a function */
func sampleFunc5() func() int {
	i := 0

	return func() int { /* an anonymous function */
		/* This function closes over the variable i
		 * to form a closure */
		i++
		return i
	}
}

func runFunctions3() {
	/* We call sampleFunc5(), assigning its result (a function)
	 * to nextInt (a function literal now). */
	nextInt := sampleFunc5()
	/* This function value captures its own i value,
	 * which will be updated each time we call nextInt. */

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt = sampleFunc5() // resets the nextInt function literal
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

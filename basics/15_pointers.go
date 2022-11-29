package basics

import "fmt"

func callByVal(arg int) {
	arg = 0
}

func callByAddr(arg *int) {
	*arg = 0
}

func runPointers() {
	i := 1
	fmt.Println("initial:", i)

	callByVal(i)
	fmt.Println("After callByVal:", i)

	callByAddr(&i)
	fmt.Println("After callByAddr:", i)

	// Print the pointer itself
	fmt.Println("pointer:", &i)
}

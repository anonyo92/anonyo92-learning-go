package basics

import (
	"fmt"
	"math"
)

const s string = "constant"

func runConstants() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d) // arbitrary precision

	fmt.Println(int64(d)) // cast

	// will infer n as float64 now (because Sin() expects that)
	fmt.Println(math.Sin(n))
}

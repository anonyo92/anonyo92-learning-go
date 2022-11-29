package basics

// import "fmt"

var p2 uint = 3

func aFuncFromAnotherSrcFile(x float64) float64 {
	// fmt.Println("Inside another function defined in a different src file of the same package")
	return x * float64(p2)
}
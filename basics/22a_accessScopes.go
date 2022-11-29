package basics

import "fmt"

var p1 int = 5 // package variable
var P3 float64 = 34.6 // exported variable (first letter capitalized); can be imported and used in other packages

func aFuncFromThisSrcFile(y float64) float64 { // package-wide visible function
	// fmt.Println("Inside another function defined in the same src file of the same package")
	return y * float64(p1)
}

func packageAccessibleFunc() { // package-wide visible function
	var a float64 = 15.7
	fmt.Println("A function is visible througout its package, but not outside its package if it is not exported.")
	fmt.Println("A function can access its parameters and variables defined locally, such as variable 'a' whose value is", a)
	fmt.Println("\t or package-variables defined outside this function in this src file, such as 'p1' whose value is", p1)
	fmt.Println("\t or package-variables defined outside this function in another src file, such as 'p2' whose value is", p2)
	fmt.Println("A function can access functions defined in the same package in the same src file, such as 'aFuncFromThisSrcFile()' which returns", aFuncFromThisSrcFile(a))
	fmt.Println("\t or functions defined in the same package in another src file, such as 'aFuncFromAnotherSrcFile()' which returns", aFuncFromAnotherSrcFile(a))
}

func RunAccessScopes() { // Exported function (first letter capitalized); can be imported and used in other packages
	packageAccessibleFunc()
	fmt.Println("A function can be exported for use in another package, by capitalizing the first letter of its name.")
}

package main

import "fmt"

func main() {

	// Initial value assigning to a variable
	var a = "initial"
	fmt.Println(a)

	// Assigning Interger to a variable
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Assigning Float to a variable
	var g, h float64 = 1.22, 22.55
	fmt.Println(g, h)

	// Assigning boolean to a variable
	var d = true
	fmt.Println(d)

	// Checking the default values
	var e int
	fmt.Println(e)

	// Shorthand to initialize and assign at the same time
	f := "Fucking shothand"
	fmt.Println(f)
}

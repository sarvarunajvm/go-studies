package main

import (
	"fmt"
	"math"
)

const str string = "This is Fucking Constant"

func main() {

	// Printing constant
	fmt.Println(str)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	// Printing the exponential values
	fmt.Println(int64(d))

	// Printing the sin value of the N
	fmt.Println(math.Sin(n))

}

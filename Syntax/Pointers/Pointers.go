package main

import "fmt"

// Value function
func zeroval(ival int) {
	ival = 0
}

// Pointer function
// /*int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences
// the pointer from its memory address to the current value
// at that address. Assigning a value to a dereferenced pointer
// changes the value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

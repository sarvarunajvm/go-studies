package main

import "fmt"

// factorial func
func fact(n int) int {
	if n == 0 {
		return 1
	}
	// Calling this func recursively
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(4))
}

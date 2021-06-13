package main

import "fmt"

// Add function with two parameters
func addTwoNumbers(a int, b int) int {
	return a + b
}

// Add function with three parameters
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

func main() {

	res := addTwoNumbers(1, 2)
	fmt.Println("1+2 =", res)

	res = addThreeNumbers(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

package main

import "fmt"

func main() {

	// Initial array with length
	var a [5]int
	fmt.Println("emp:", a)

	// Setting and getting the value in array
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Printing the length using default function
	fmt.Println("len:", len(a))

	// Declaring the array with shorthand
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Two dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

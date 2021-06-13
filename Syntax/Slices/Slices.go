package main

import "fmt"

// Unlike arrays, slices are typed only by the elements
// they contain (not the number of elements). To create
// an empty slice with non-zero length, use the builtin
// make. Here we make a slice of strings of length 3
// (initially zero-valued).
func main() {

	// Initial slices with length 3
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// setting and getting values in the slices as same as array
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Printing the length using default function
	fmt.Println("len:", len(s))

	// Appending into slices
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy the slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// cutting the slices from 2-5
	l := s[2:5]
	fmt.Println("sl1:", l)

	// cutting the slices from 0-5
	l = s[:5]
	fmt.Println("sl2:", l)

	// cutting the slices from 2-5
	l = s[2:]
	fmt.Println("sl3:", l)

	// Declaring the slices with shorthand
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// TwoDimensional slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

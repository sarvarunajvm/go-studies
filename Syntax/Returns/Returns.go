package main

import "fmt"

// Returning multiple values from single function
func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

	// If you only want a subset 
	// of the returned values, use the blank identifier _.
	// with blank identifiers
    _, c := vals()
    fmt.Println(c)
}
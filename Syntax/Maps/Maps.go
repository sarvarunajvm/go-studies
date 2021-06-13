package main

import "fmt"

func main() {

	// Declaring the maps with shorthand
	m := make(map[string]int)

	// setting the value in map
	m["k1"] = 7
	m["k2"] = 13

	// printing the Map
	fmt.Println("map:", m)

	// getting the value from map
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Printing the length using default function
	fmt.Println("len:", len(m))

	// Deleting using the key
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting
	// a value from a map indicates if the key was
	// present in the map. This can be used to disambiguate
	// between missing keys and keys with zero values like 0
	// or "". Here we didn’t need the value itself, so
	// we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declaring the maps with values using the shorthand
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

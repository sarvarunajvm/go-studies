package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	// For Each also called Range used to iterates over elements
	// With blank identifier
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// With temp assigning variable
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Over the key value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Over the Keys from the map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Over the Unicode points
	for i, c := range "FUCK GO" {
		fmt.Println(i, c)
	}
}

package main

import "fmt"

// Arbitary number of arguments 
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

	// Below are temp array
    sum(1, 2)
    sum(1, 2, 3)

	// New array or slices
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
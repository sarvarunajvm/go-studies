package main

import "fmt"

// Closures and also called as inline functions
// Basically it returns a func declared in the body of the another func
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

	// Assinging the closure to variable 
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

	// Re-Assinging the closure to variable 
	// below should print from start
    newInts := intSeq()
    fmt.Println(newInts())
}
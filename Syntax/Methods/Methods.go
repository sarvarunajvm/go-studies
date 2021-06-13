package main

import "fmt"

type rect struct {
    width, height int
}

// Function to find area 
//This area method has a receiver type of *rect.
func (r *rect) area() int {
    return r.width * r.height
}

// Function to find perimeter 
func (r rect) perimeter() int {
    return 2*r.width + 2*r.height
}

// It automatically handles conversion between values and pointers 
// for method calls. You may want to use a pointer receiver type 
// to avoid copying on method calls or to allow the method to 
// mutate the receiving struct.
func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("Area: ", r.area())
    fmt.Println("Perimeter:", r.perimeter())

    rp := &r
    fmt.Println("Area: ", rp.area())
    fmt.Println("Perimeter:", rp.perimeter())
}
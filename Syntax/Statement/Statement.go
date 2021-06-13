package main

import "fmt"

func main() {

    // Basic If and else
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // Only if without an else
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // ElseIf's 
    if num := 11; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}